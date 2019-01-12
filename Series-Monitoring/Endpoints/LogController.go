package Endpoints

import (
	"net/http"
	"encoding/json"
	"github.com/fgehrlicher/series-monitoring/Series-Monitoring/Models"
	"strconv"
)

func GetAllLogs(response http.ResponseWriter, request *http.Request) {
	var logs []Models.Log
	var err error
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()

	queryParameter := request.URL.Query()
	logRepository := Models.LogRepository{Db: database}
	since, passed := queryParameter["since"]

	if passed {
		since, _ := strconv.Atoi(since[0])
		logs, err = logRepository.GetAllSince(since)
	} else {
		logs, err = logRepository.GetAll()
	}

	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	json.NewEncoder(response).Encode(logs)
	logAccess(database, request)
}

func GetAllMessages(response http.ResponseWriter, request *http.Request) {
	GetAllByType(response, request, Models.LogTypeMessage)
}

func GetAllWarnings(response http.ResponseWriter, request *http.Request) {
	GetAllByType(response, request, Models.LogTypeWarning)

}

func GetAllErrors(response http.ResponseWriter, request *http.Request) {
	GetAllByType(response, request, Models.LogTypeError)
}

func GetAllByType(response http.ResponseWriter, request *http.Request, logType int) {
	var logs []Models.Log
	var err error
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()

	queryParameter := request.URL.Query()
	logRepository := Models.LogRepository{Db: database}
	since, passed := queryParameter["since"]

	if passed {
		since, _ := strconv.Atoi(since[0])
		logs, err = logRepository.GetAllSinceByType(since, logType)
	} else {
		logs, err = logRepository.GetAllByType(logType)
	}

	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	json.NewEncoder(response).Encode(logs)
	logAccess(database, request)
}
