package Endpoints

import (
	"net/http"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
)

/*
[
  {
    "ImageID": 5,
    "ImagePath": "Provider\/netflix.png",
    "Name": "Netflix",
    "ID": 1
  },
  {
    "ImageID": 4,
    "ImagePath": "Provider\/burningseries.png",
    "Name": "Burning Series",
    "ID": 2
  },
  {
    "ImageID": 6,
    "ImagePath": "Provider\/primevideo.png",
    "Name": "Amazon Prime",
    "ID": 3
  }
]
 */
func GetAllProviders(response http.ResponseWriter, request *http.Request) {
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	providerRepository := Models.ProviderRepository{Db: database}
	providers, err := providerRepository.GetAll()
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	json.NewEncoder(response).Encode(providers)
	logAccess(database, request)
}

/*
{
  "ImageID": 5,
  "ImagePath": "Provider\/netflix.png",
  "Name": "Netflix",
  "ID": 1
}
 */
func GetProvider(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	providerName := vars["provider"]
	if providerName == "" {
		NotFoundHandler(response, request)
		return
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	providerRepository := Models.ProviderRepository{Db: database}
	provider, err := providerRepository.GetByName(providerName)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}
	json.NewEncoder(response).Encode(provider)
	logAccess(database, request)
}

/*
Png Image
 */
func GetProviderImage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	providerName := vars["provider"]
	if providerName == "" {
		NotFoundHandler(response, request)
		return
	}
	settings, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	providerRepository := Models.ProviderRepository{Db: database}
	provider, err := providerRepository.GetByName(providerName)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}
	image := Models.Image{RelativePath: provider.ImagePath, Settings: settings}
	err = image.LoadImageFromFile(Models.ImageProvider)
	if err != nil {
		NotFoundHandler(response, request)
		return
	}
	//@TODO other content types
	response.Header().Set("Content-type", "image/png")
	response.Write(image.Data)
	logAccess(database, request)
}
