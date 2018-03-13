package Endpoints

import (
	"net/http"
	"encoding/json"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"strconv"
)

/*
[
  {
    "Message": "Authorized Call",
    "Time": "2018-03-12T07:24:21Z",
    "Type": 1,
    "Caller": "127.0.0.1:53419",
    "RequestUri": "\/logs\/",
    "StatusCode": 200,
    "ID": 0
  },
  ....
]
 */
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

/*
[
  {
    "Message": "Authorized Call",
    "Time": "2018-03-12T07:24:21Z",
    "Type": 1,
    "Caller": "127.0.0.1:53419",
    "RequestUri": "\/logs\/",
    "StatusCode": 200,
    "ID": 0
  },
  ....
]
 */
func GetAllMessages(response http.ResponseWriter, request *http.Request) {
	GetAllByType(response, request, Models.LogTypeMessage)
}

/*
[
  {
    "Message": "Not Found",
    "Time": "2018-03-09T13:51:01Z",
    "Type": 2,
    "Caller": "127.0.0.1:58052",
    "RequestUri": "\/logs\/adasd",
    "StatusCode": 404,
    "ID": 0
  },
  ....
]
 */
func GetAllWarnings(response http.ResponseWriter, request *http.Request) {
	GetAllByType(response, request, Models.LogTypeWarning)

}

/*
[
  {
    "Message": "Server Error",
    "Time": "2018-03-09T13:51:01Z",
    "Type": 3,
    "Caller": "127.0.0.1:58052",
    "RequestUri": "\/logs\/",
    "StatusCode": 404,
    "ID": 0
  },
  ....
]
 */
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
