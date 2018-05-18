package Endpoints

import (
	"net/http"
	"github.com/gorilla/mux"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"database/sql"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Parser"
	"strconv"
	"github.com/fatih/color"
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

	err = updateSeries(series, database)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}
	database.Close()
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

	for _, series := range seriesSlice {
		err := updateSeries(&series, database)
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
	}
}

func updateSeries(series *Models.Series, database *sql.DB) error {
	parser := Parser.TMDbHandler{Series: series, SeriesUrl: series.ProviderURL}
	episodeRepository := Models.EpisodeRepository{Db: database}
	imageRepository := Models.ImageRepository{Db: database}
	logRepository := Models.LogRepository{Db: database}

	latestEpisode, err := episodeRepository.GetLatestBySeries(*series)
	if err != nil {
		if err == sql.ErrNoRows {
			latestEpisode = &Models.Episode{Season: 1, Episode: 0}
		} else {
			return err
		}
	}
	episodes, err := parser.GetAllNewEpisodes(*latestEpisode)

	if err != nil {
		return err
	}
	if len(episodes) == 0 {
		return nil
	}
	for _, episode := range episodes {
		if episode.Image.OriginURL != "" {
			imageId, err := imageRepository.Persist(*episode.Image)
			if err != nil {
				return err
			}
			episode.Image.ID = imageId
		}
		episode.Series = series
		_, err := episodeRepository.Persist(episode)
		if err != nil {
			color.Red(err.Error() + "\n")
		}
	}

	logRepository.Persist(
		Models.Log{
			Type:       Models.LogTypeMessage,
			Caller:     "-",
			RequestUri: "-",
			StatusCode: 200,
			Message:    "Imported " + strconv.Itoa(len(episodes)) + " Episode for Series: " + series.Title,
		},
	)
	return nil
}

func checkImageIntegrity(response http.ResponseWriter, request *http.Request) {

}