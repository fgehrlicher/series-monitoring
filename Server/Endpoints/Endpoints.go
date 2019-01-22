package Endpoints

import (
	"net/http"
	"github.com/fgehrlicher/series-monitoring/Server/Config"
	"database/sql"
	"github.com/fgehrlicher/series-monitoring/Server/Models"
	"github.com/fatih/color"
	"fmt"
	"os"
	"encoding/json"
)

type Message struct {
	Message string `json:"message"`
}

func NewSuccessMessage() Message {
	return Message{Message: "Success"}
}
func NewFailureMessage() Message {
	return Message{Message: "failure"}
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(Message{"Series watcher v3"})
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
