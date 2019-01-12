package Endpoints

import (
	"net/http"
	"github.com/fgehrlicher/series-monitoring/Series-Monitoring/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
)

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
