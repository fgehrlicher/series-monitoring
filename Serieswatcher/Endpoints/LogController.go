package Endpoints

import (
	"net/http"
	"encoding/json"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"github.com/gorilla/mux"
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
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAll()
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
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAllByType(Models.LogTypeMessage)
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
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAllByType(Models.LogTypeWarning)
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
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAllByType(Models.LogTypeError)
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
func GetLogsSince(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	since, err := strconv.Atoi(vars["since"])
	if err != nil || since == 0 {
		GetAllLogs(response, request)
		return
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAllSince(since)
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
func GetMessagesSince(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	since, err := strconv.Atoi(vars["since"])
	if err != nil || since == 0 {
		GetAllLogs(response, request)
		return
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAllSinceByType(since, Models.LogTypeWarning)
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
func GetWarningsSince(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	since, err := strconv.Atoi(vars["since"])
	if err != nil || since == 0 {
		GetAllLogs(response, request)
		return
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAllSinceByType(since, Models.LogTypeWarning)
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
func GetErrorsSince(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	since, err := strconv.Atoi(vars["since"])
	if err != nil || since == 0 {
		GetAllLogs(response, request)
		return
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	logRepository := Models.LogRepository{Db: database}
	logs, err := logRepository.GetAllSinceByType(since, Models.LogTypeError)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	json.NewEncoder(response).Encode(logs)
	logAccess(database, request)
}
