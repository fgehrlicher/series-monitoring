package Server

import (
	"../Config"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/fatih/color"
)

var settings Config.Settings

func Init() {

	var err error
	settings, err = Config.GetConfiguration()
	if err != nil {
		panic(err.Error())
	}
	router := mux.NewRouter()
	ip := settings.ServerSettings.Ip
	port := settings.ServerSettings.Port

	router.StrictSlash(true)
	AttachEndpoints(router)
	color.Green("Server listening on: " + ip + ":" + port)
	panic(http.ListenAndServe(ip+":"+port, router))
}

func AttachEndpoints(router *mux.Router) {

	/**
	router.HandleFunc("/functions/update-series/{series-name}", updateSeries).Methods("POST")
	router.HandleFunc("/functions/update-series/*", updateAllSeries).Methods("POST")
	router.HandleFunc("/functions/process-queue/", processQueue).Methods("POST")

	router.HandleFunc("/provider/", getAllProvider).Methods("GET")
	router.HandleFunc("/provider/{provider}/image", getProviderImage).Methods("GET")

	router.HandleFunc("/series/", getAllSeries).Methods("GET")
	router.HandleFunc("/series/{series}/", getSeries).Methods("GET")
	router.HandleFunc("/series/{series}/", createSeries).Methods("POST")
	router.HandleFunc("/series/{series}/", updateSeries).Methods("PUT")
	router.HandleFunc("/series/{series}/", deleteSeries).Methods("DELETE")
	router.HandleFunc("/series/{series}/pointer", deleteSeries).Methods("GET")
	router.HandleFunc("/series/{series}/pointer", movePointer).Methods("POST")
	router.HandleFunc("/series/{series}/image", getSeriesImage).Methods("GET")
	router.HandleFunc("/series/{series}/unwatched-episode/", getUnwatchedEpisodes).Methods("GET")
	router.HandleFunc("/series/{series}/episode/{season}/{episode}/", getEpisode).Methods("GET")
	router.HandleFunc("/series/{series}/episode/{season}/{episode}/image", getEpisodeImage).Methods("GET")
	*/


}
