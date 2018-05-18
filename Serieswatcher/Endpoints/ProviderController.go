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
	providers, err := providerRepository.GetAll(false)
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
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	providerRepository := Models.ProviderRepository{Db: database}
	provider, err := providerRepository.GetByName(providerName, false)
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
png/jpg Image
 */
func GetProviderImage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	providerName := vars["provider"]
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	providerRepository := Models.ProviderRepository{Db: database}
	provider, err := providerRepository.GetByName(providerName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}
	if !(provider.Image.ID > 0) {
		NotFoundHandler(response, request)
		return
	}
	err = provider.Image.LoadFromFile()
	if err != nil {
		NotFoundHandler(response, request)
		return
	}
	response.Write(provider.Image.Data)
	logAccess(database, request)
}
