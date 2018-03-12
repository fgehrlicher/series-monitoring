package Endpoints

import (
	"github.com/gorilla/mux"
	"net/http"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Config"
	"database/sql"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"github.com/fatih/color"
	"fmt"
	"os"
	"encoding/json"
)

type Message struct {
	Message string `json:"message"`
}

/*
Returns:

{
	"message": "Series watcher v3"
}
 */
func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(Message{"Series watcher v3"})
}

func AttachEndpoints(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	router.HandleFunc("/", RootEndpoint).Methods("GET")

	router.HandleFunc("/logs/", GetAllLogs).Methods("GET")
	router.HandleFunc("/logs/{since:[0-9]+}", GetLogsSince).Methods("GET")
	router.HandleFunc("/logs/message/", GetAllMessages).Methods("GET")
	router.HandleFunc("/logs/warning/", GetAllWarnings).Methods("GET")
	router.HandleFunc("/logs/error/", GetAllErrors).Methods("GET")
	router.HandleFunc("/logs/message/{since:[0-9]+}", GetMessagesSince).Methods("GET")
	router.HandleFunc("/logs/warning/{since:[0-9]+}", GetWarningsSince).Methods("GET")
	router.HandleFunc("/logs/error/{since:[0-9]+}", GetErrorsSince).Methods("GET")

	router.HandleFunc("/provider/", GetAllProvider).Methods("GET")
	router.HandleFunc("/provider/{provider}", GetProvider).Methods("GET")
	router.HandleFunc("/provider/{provider}/image", GetProviderImage).Methods("GET")

	/*
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

		router.HandleFunc("/functions/update-series/{series-name}", updateSeries).Methods("POST")
		router.HandleFunc("/functions/update-series/*", updateAllSeries).Methods("POST")
		router.HandleFunc("/functions/process-queue/", processQueue).Methods("POST")
		*/

}

func logAccess(database *sql.DB, request *http.Request) {
	repository := Models.LogRepository{Db: database}

	log := Models.Log{
		Message:    "Authorized Call",
		Type:       Models.LogTypeMessage,
		Caller:     request.RemoteAddr,
		RequestUri: request.RequestURI,
		StatusCode: 200,
	}

	if err := repository.Persist(log); err != nil {
		fmt.Fprintln(os.Stderr, color.RedString("[ERROR STACK] "+err.Error()+"\n"))
	}
}

func getSettingsAndDatabase(response http.ResponseWriter, request *http.Request) (settings Config.Settings, db *sql.DB) {
	settings, err := Config.GetConfiguration()
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}

	db, err = Models.GetDatabaseConnection(settings)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}

	return settings, db
}
