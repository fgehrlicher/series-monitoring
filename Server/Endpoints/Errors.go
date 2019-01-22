package Endpoints

import (
	"net/http"
	"github.com/fgehrlicher/series-monitoring/Server/Models"
	"encoding/json"
	"errors"
	"github.com/fgehrlicher/series-monitoring/Server/Config"
	"github.com/fatih/color"
	"database/sql"
	"fmt"
	"os"
)

/*
Returns:

{
   "message": "Not Found"
}
 */
func NotFoundHandler(responseWriter http.ResponseWriter, request *http.Request) {
	handleError(
		responseWriter,
		request,
		errors.New(http.StatusText(http.StatusNotFound)),
		http.StatusNotFound,
		Models.LogTypeWarning,
	)
}

/*
Returns:

{
   "message": "Custom Error Message"
}
 */
func NotFoundErrorHandler(responseWriter http.ResponseWriter, request *http.Request, err error) {
	handleError(
		responseWriter,
		request,
		err,
		http.StatusNotFound,
		Models.LogTypeWarning,
	)
}

/*
Returns:

{
   "message": "Not Found"
}
 */
func BadRequestHandler(responseWriter http.ResponseWriter, request *http.Request) {
	handleError(
		responseWriter,
		request,
		errors.New(http.StatusText(http.StatusBadRequest)),
		http.StatusBadRequest,
		Models.LogTypeWarning,
	)
}

/*
Returns:

{
   "message": "Method Not Allowed"
}
 */
func MethodNotAllowedHandler(responseWriter http.ResponseWriter, request *http.Request) {
	handleError(
		responseWriter,
		request,
		errors.New(http.StatusText(http.StatusMethodNotAllowed)),
		http.StatusMethodNotAllowed,
		Models.LogTypeWarning,
	)
}

/*
Returns:

{
   "message": "Internal Server Error"
}
 */
func InternalServerErrorHandler(responseWriter http.ResponseWriter, request *http.Request, err error) {
	fmt.Fprintln(os.Stderr, color.RedString("[ERROR STACK] "+err.Error()+"\n"))
	handleError(
		responseWriter,
		request,
		errors.New(http.StatusText(http.StatusInternalServerError)+" "+err.Error()),
		http.StatusInternalServerError,
		Models.LogTypeError,
	)
}

func handleError(response http.ResponseWriter, request *http.Request, err error, errorCode int, logType int) {
	var (
		errorMessage string
		internalErr  error
		settings     Config.Settings
		database     *sql.DB
		repository   Models.LogRepository
	)

	errorMessage = err.Error()
	settings, internalErr = Config.GetConfiguration()
	if internalErr != nil {
		fmt.Fprintln(os.Stderr, color.RedString("[ERROR STACK] "+err.Error()+"\n"))
	}

	database, internalErr = Models.GetDatabaseConnection(settings)
	if internalErr != nil {
		fmt.Fprintln(os.Stderr, color.RedString("[ERROR STACK] "+err.Error()+"\n"))
	}
	defer database.Close()

	response.WriteHeader(errorCode)
	json.NewEncoder(response).Encode(Message{errorMessage})

	log := Models.Log{
		Message:    errorMessage,
		Type:       logType,
		Caller:     request.RemoteAddr,
		RequestUri: request.RequestURI,
		StatusCode: errorCode,
	}

	repository = Models.LogRepository{Db: database}

	if internalErr = repository.Persist(log); internalErr != nil {
		fmt.Fprintln(os.Stderr, color.RedString("[ERROR STACK] "+err.Error()+"\n"))
	}
}
