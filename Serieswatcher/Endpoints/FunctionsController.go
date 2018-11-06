package Endpoints

import (
	"net/http"
	"github.com/gorilla/mux"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"database/sql"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Parser"
	"strconv"
	"github.com/fatih/color"
	"encoding/json"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Maintenance"
)

func UpdateSeries(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()

	seriesRepository := Models.SeriesRepository{Db: database}

	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	message, err := updateSeries(series, database)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}

	if message == "" {
		message = "no new episodes found for series: " + series.Title
	}

	json.NewEncoder(response).Encode(
		struct {
			Message string
			Steps   []Message
		}{
			Message: "Success",
			Steps:   []Message{{Message: message}},
		})
	logAccess(database, request)
}

func UpdateAllSeries(response http.ResponseWriter, request *http.Request) {
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()

	seriesRepository := Models.SeriesRepository{Db: database}

	seriesSlice, err := seriesRepository.GetAll(true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	messages := make([]Message, 0)
	for _, series := range seriesSlice {
		message, err := updateSeries(&series, database)
		if message != "" {
			messages = append(messages, Message{Message:message})
		}
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
	}

	if len(messages) == 0 {
		messages = append(messages, Message{Message: "no new episodes found"})
	}

	json.NewEncoder(response).Encode(
		struct {
			Message string
			Steps   []Message
		}{
			Message: "Success",
			Steps:   messages,
		})
	logAccess(database, request)
}

func updateSeries(series *Models.Series, database *sql.DB) (string, error) {
	parser := Parser.TMDbHandler{Series: series, SeriesUrl: series.ProviderURL}
	episodeRepository := Models.EpisodeRepository{Db: database}
	imageRepository := Models.ImageRepository{Db: database}
	logRepository := Models.LogRepository{Db: database}

	var message string
	latestEpisode, err := episodeRepository.GetLatestBySeries(*series)
	if err != nil {
		if err == sql.ErrNoRows {
			latestEpisode = &Models.Episode{Season: 1, Episode: 0}
		} else {
			return message, err
		}
	}
	episodes, err := parser.GetAllNewEpisodes(*latestEpisode)

	if err != nil {
		return message, err
	}
	if len(episodes) == 0 {
		return message, err
	}
	for _, episode := range episodes {
		if episode.Image.OriginURL != "" {
			imageId, err := imageRepository.Persist(*episode.Image)
			if err != nil {
				return message, err
			}
			episode.Image.ID = imageId
		}
		episode.Series = series
		_, err := episodeRepository.Persist(episode)
		if err != nil {
			color.Red(err.Error() + "\n")
		}
	}

	message = "Imported " + strconv.Itoa(len(episodes)) + " Episodes for Series: " + series.Title
	logRepository.Persist(
		Models.Log{
			Type:       Models.LogTypeMessage,
			Caller:     "-",
			RequestUri: "-",
			StatusCode: 200,
			Message:    message,
		},
	)
	return message, nil
}

func CheckIntegrity(response http.ResponseWriter, request *http.Request) {
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()

	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	imageRepository := Models.ImageRepository{Db: database}
	logRepository := Models.LogRepository{Db: database}
	integrityController := Maintenance.NewIntegrityController()

	images, err := imageRepository.GetAll()
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	downloadCount, message := integrityController.CheckImageIntegrity(images)
	logRepository.Persist(
		Models.Log{
			Type:       Models.LogTypeMessage,
			Caller:     "-",
			RequestUri: "-",
			StatusCode: 200,
			Message:    "Downloaded " + strconv.Itoa(downloadCount) + " Images. \n Raw Message: \n" + message,
		},
	)

	seriesSlice, err := seriesRepository.GetAll(true)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}

	episodesToCheck := make([]Models.Episode, 0)

	for _, series := range seriesSlice {
		episodes, err := episodeRepository.GetAllNewEpisodes(series)
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
		if len(episodes) > 0 {
			episodesToCheck = append(episodesToCheck, episodes...)
		}
	}

	updateCount, message := integrityController.CheckMetaDataIntegrity(episodesToCheck)
	logRepository.Persist(
		Models.Log{
			Type:       Models.LogTypeMessage,
			Caller:     "-",
			RequestUri: "-",
			StatusCode: 200,
			Message:    "Updated " + strconv.Itoa(updateCount) + " Episodes. \n Raw Message: \n" + message,
		},
	)

	logAccess(database, request)
}
