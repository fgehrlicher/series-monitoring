package Server

import (
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Config"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/fatih/color"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Endpoints"
	"time"
	"os"
	"os/signal"
	"context"
)

func Init() {
	settings, err := Config.GetConfiguration()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	router := mux.NewRouter()
	ip := settings.ServerSettings.Ip
	port := settings.ServerSettings.Port

	router.StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(Endpoints.NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(Endpoints.MethodNotAllowedHandler)

	router.HandleFunc("/", Endpoints.RootEndpoint).Methods("GET")

	router.HandleFunc("/logs/", Endpoints.GetAllLogs).Methods("GET")
	router.HandleFunc("/logs/message/", Endpoints.GetAllMessages).Methods("GET")
	router.HandleFunc("/logs/warning/", Endpoints.GetAllWarnings).Methods("GET")
	router.HandleFunc("/logs/error/", Endpoints.GetAllErrors).Methods("GET")

	router.HandleFunc("/provider/", Endpoints.GetAllProviders).Methods("GET")
	router.HandleFunc("/provider/{provider}/", Endpoints.GetProvider).Methods("GET")
	router.HandleFunc("/provider/{provider}/image", Endpoints.GetProviderImage).Methods("GET")

	router.HandleFunc("/series/", Endpoints.GetAllSeries).Methods("GET")
	router.HandleFunc("/series/{series}/", Endpoints.GetSeries).Methods("GET")
	router.HandleFunc("/series/{series}/image", Endpoints.GetSeriesImage).Methods("GET")
	router.HandleFunc("/series/{series}/episodes", Endpoints.GetAllEpisodes).Methods("GET")
	router.HandleFunc("/series/{series}/episodes/unwatched", Endpoints.GetNewEpisodes).Methods("GET")
	router.HandleFunc("/series/{series}/season/{season}", Endpoints.GetAllEpisodesBySeason).Methods("GET")
	router.HandleFunc("/series/{series}/season/{season}/episode/{episode}", Endpoints.GetEpisode).Methods("GET")
	router.HandleFunc("/series/{series}/season/{season}/episode/{episode}/image", Endpoints.GetEpisodeImage).Methods("GET")

	router.HandleFunc("/series/", Endpoints.CreateSeries).Methods("POST")
	router.HandleFunc("/series/{series}/pointer", Endpoints.MovePointer).Methods("POST")
	router.HandleFunc("/series/{series}/update", Endpoints.UpdateSeries).Methods("POST")

	router.HandleFunc("/functions/update-all-series/", Endpoints.UpdateAllSeries).Methods("POST")

	server := &http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	color.Green("Server listening on: " + ip + ":" + port)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			server.ListenAndServe()
		}
	}()

	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt)

	<-osSignalChannel

	ctx, cancelFunction := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunction()
	server.Shutdown(ctx)
	color.Green("shutting down")
	os.Exit(0)

}
