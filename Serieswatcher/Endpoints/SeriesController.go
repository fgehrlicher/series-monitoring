package Endpoints

import (
	"net/http"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Parser"
	"strings"
)

/*
[
  {
    "ID": 4,
    "ImageID": 4,
    "ImagePath": "Series\/Bojack-horseman.png",
    "ProviderURL": "tv/2424323",
    "Title": "Bojack Horeseman"
  },
  ....
]
 */
func GetAllSeries(response http.ResponseWriter, request *http.Request) {
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	series, err := seriesRepository.GetAll()
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	json.NewEncoder(response).Encode(series)
	logAccess(database, request)
}

/*
{
  "ID": 4,
  "ImageID": 4,
  "ImagePath": "Series\/Bojack-horseman.png",
  "ProviderURL": "tv/2424323",
  "Title": "Bojack Horeseman"
}
 */
func GetSeries(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}
	json.NewEncoder(response).Encode(series)
	logAccess(database, request)
}

/*
png/jpg image
 */
func GetSeriesImage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	settings, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}
	image := Models.Image{RelativePath: series.ImagePath, Settings: settings}
	err = image.LoadImageFromFile(Models.ImageProvider)
	if err != nil {
		NotFoundHandler(response, request)
		return
	}
	response.Write(image.Data)
	logAccess(database, request)
}

/*
{
  "ID": 4,
  "ImageID": 4,
  "ImagePath": "Series\/Bojack-horseman.png",
  "ProviderURL": "tv/2424323",
  "Title": "Bojack Horeseman"
}
 */
func CreateSeries(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	seriesUrl, passed := request.Form["seriesUrl"]
	if !passed || !strings.HasPrefix(seriesUrl[0], Parser.TMDbBaseUrl){
		BadRequestHandler(response, request)
		return
	}
	settings, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}

	_, err := seriesRepository.GetByProviderURL(seriesUrl[0])
	if err == nil {
		BadRequestHandler(response, request)
		return
	}
	handler, err := Parser.NewTMDbHandler(seriesUrl[0], settings)
	if err != nil {
		NotFoundErrorHandler(response, request, err)
		return
	}
	series, err := handler.GetSeries()
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	err = seriesRepository.Persist(series)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	newSeries, err := seriesRepository.GetByName(series.Title)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	json.NewEncoder(response).Encode(newSeries)
	logAccess(database, request)
}
