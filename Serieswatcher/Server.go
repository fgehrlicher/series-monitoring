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
		panic(err.Error())
	}

	router := mux.NewRouter()
	ip := settings.ServerSettings.Ip
	port := settings.ServerSettings.Port

	router.StrictSlash(true)
	Endpoints.AttachEndpoints(router)

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
