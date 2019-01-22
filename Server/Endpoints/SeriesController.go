package Endpoints

import (
	"net/http"
	"github.com/fgehrlicher/series-monitoring/Server/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
	"github.com/fgehrlicher/series-monitoring/Server/Parser"
	"strings"
	"strconv"
	"errors"
)

func GetAllSeries(response http.ResponseWriter, request *http.Request) {
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	series, err := seriesRepository.GetAll(true)
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
	series, err := seriesRepository.GetByName(seriesName, true)
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

	err = series.Image.LoadFromFile()
	if err != nil {
		NotFoundHandler(response, request)
		return
	}

	key := "black"
	e := `"` + key + `"`
	response.Header().Set("Etag", e)
	response.Header().Set("Cache-Control", "max-age=2592000")
	response.Write(series.Image.Data)
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
	seriesUrlSlice, passed := request.Form["series_url"]
	if !passed {
		BadRequestHandler(response, request)
		return
	}
	seriesUrl := seriesUrlSlice[0]

	if !passed || !strings.HasPrefix(seriesUrl, Parser.TMDbBaseUrl) {
		BadRequestHandler(response, request)
		return
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	imageRepository := Models.ImageRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}

	_, err := seriesRepository.GetByProviderURL(seriesUrl, false)
	if err != sql.ErrNoRows {
		BadRequestHandler(response, request)
		return
	}

	handler, err := Parser.NewTMDbHandler(Models.Series{ProviderURL: seriesUrl})
	if err != nil {
		NotFoundErrorHandler(response, request, err)
		return
	}
	series, err := handler.GetSeries()
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}

	if series.Image.RelativePath != "" {
		imageID, err := imageRepository.Persist(*series.Image)
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
		series.Image.ID = imageID
	}

	seriesID, err := seriesRepository.Persist(*series)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	series.ID = seriesID
	series.WatchPointer.Series = &Models.Series{ID: seriesID}

	if series.WatchPointer.Image.ID == 0 {
		watchpointerImageID, err := imageRepository.Persist(*series.WatchPointer.Image)
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
		series.WatchPointer.Image.ID = watchpointerImageID
	}

	episodeID, err := episodeRepository.Persist(*series.WatchPointer)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	series.WatchPointer.ID = episodeID

	seriesRepository.UpdateWatchPointer(series)

	json.NewEncoder(response).Encode(series)
	logAccess(database, request)
}

/*
[
	{
		"ID": 391,
		"Episode": 1,
		"Season": 1,
		"Title": "The Bone Orchard",
		"Description": "When Shadow Moon is released from prison early after the death of his wife, he meets Mr. Wednesday and is recruited as his bodyguard. Shadow discovers that this may be more than he bargained for.",
		"ReleaseDate": "2017-04-30T00:00:00Z"
	},
	...
]
*/
func GetAllEpisodes(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	episodes, err := episodeRepository.GetAllBySeries(*series, true)
	json.NewEncoder(response).Encode(episodes)
	logAccess(database, request)
}

/*
[
	{
		"ID": 391,
		"Episode": 1,
		"Season": 1,
		"Title": "The Bone Orchard",
		"Description": "When Shadow Moon is released from prison early after the death of his wife, he meets Mr. Wednesday and is recruited as his bodyguard. Shadow discovers that this may be more than he bargained for.",
		"ReleaseDate": "2017-04-30T00:00:00Z"
	},
	...
]
*/
func GetAllEpisodesBySeason(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	seasonString := vars["season"]
	season, err := strconv.Atoi(seasonString)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	episodes, err := episodeRepository.GetAllBySeriesAndSeason(*series, season, true)
	json.NewEncoder(response).Encode(episodes)
	logAccess(database, request)
}

/*
{
	"ID": 391,
	"Episode": 1,
	"Season": 1,
	"Title": "The Bone Orchard",
	"Description": "When Shadow Moon is released from prison early after the death of his wife, he meets Mr. Wednesday and is recruited as his bodyguard. Shadow discovers that this may be more than he bargained for.",
	"ReleaseDate": "2017-04-30T00:00:00Z"
}
*/
func GetEpisode(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	seasonString := vars["season"]
	episodeString := vars["episode"]
	season, err := strconv.Atoi(seasonString)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	episode, err := strconv.Atoi(episodeString)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	element, err := episodeRepository.GetOneBySeriesAndSeasonAndEpisode(*series, season, episode, true)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}
	json.NewEncoder(response).Encode(element)
	logAccess(database, request)
}

/*
png/jpg image
*/
func GetEpisodeImage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	seasonString := vars["season"]
	episodeString := vars["episode"]
	season, err := strconv.Atoi(seasonString)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}
	episode, err := strconv.Atoi(episodeString)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	element, err := episodeRepository.GetOneBySeriesAndSeasonAndEpisode(*series, season, episode, true)
	if err != nil {
		NotFoundHandler(response, request)
		return
	}

	if element.Image != nil {
		err = element.Image.LoadFromFile()
		if err != nil {
			NotFoundHandler(response, request)
			return
		}
		key := "black"
		e := `"` + key + `"`
		response.Header().Set("Etag", e)
		response.Header().Set("Cache-Control", "max-age=2592000")
		response.Write(element.Image.Data)
	} else {
		image := Models.Image{
			RelativePath: "none",
			ImageType:    Models.ImageEpisode,
		}
		err = image.LoadFromFile()
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
		response.Write(image.Data)
	}

	logAccess(database, request)
}

/*
[
	{
		"ID": 391,
		"Episode": 1,
		"Season": 1,
		"Title": "The Bone Orchard",
		"Description": "When Shadow Moon is released from prison early after the death of his wife, he meets Mr. Wednesday and is recruited as his bodyguard. Shadow discovers that this may be more than he bargained for.",
		"ReleaseDate": "2017-04-30T00:00:00Z"
	},
	...
]
*/
func GetNewEpisodes(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]

	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	episodes, err := episodeRepository.GetAllNewEpisodes(*series)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}
	json.NewEncoder(response).Encode(episodes)
	logAccess(database, request)
}

/*
[
	{
		"ID": 391,
		"Episode": 1,
		"Season": 1,
		"Title": "The Bone Orchard",
		"Description": "When Shadow Moon is released from prison early after the death of his wife, he meets Mr. Wednesday and is recruited as his bodyguard. Shadow discovers that this may be more than he bargained for.",
		"ReleaseDate": "2017-04-30T00:00:00Z"
	},
	...
]
*/
func GetUpcomingEpisodes(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]

	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	episodes, err := episodeRepository.GetAllUpcomingEpisodes(*series)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}
	json.NewEncoder(response).Encode(episodes)
	logAccess(database, request)
}

/*
[
	{
		"ID": 391,
		"Episode": 1,
		"Season": 1,
		"Title": "The Bone Orchard",
		"Description": "When Shadow Moon is released from prison early after the death of his wife, he meets Mr. Wednesday and is recruited as his bodyguard. Shadow discovers that this may be more than he bargained for.",
		"ReleaseDate": "2017-04-30T00:00:00Z"
	},
	...
]
*/
func GetAllNewEpisodes(response http.ResponseWriter, request *http.Request) {
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	seriesSlice, err := seriesRepository.GetAll(true)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}

	resultSlice := make([]Models.Series, 0)

	for _, series := range seriesSlice {
		episodes, err := episodeRepository.GetAllNewEpisodes(series)
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
		if len(episodes) > 0 {
			series.UnwatchedEpisodes = episodes
			resultSlice = append(resultSlice, series)
		}
	}

	json.NewEncoder(response).Encode(resultSlice)
	logAccess(database, request)
}

/*
[
	{
		"ID": 391,
		"Episode": 1,
		"Season": 1,
		"Title": "The Bone Orchard",
		"Description": "When Shadow Moon is released from prison early after the death of his wife, he meets Mr. Wednesday and is recruited as his bodyguard. Shadow discovers that this may be more than he bargained for.",
		"ReleaseDate": "2017-04-30T00:00:00Z"
	},
	...
]
*/
func GetAllUpcomingEpisodes(response http.ResponseWriter, request *http.Request) {
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}
	seriesSlice, err := seriesRepository.GetAll(true)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}

	resultSlice := make([]Models.Series, 0)

	for _, series := range seriesSlice {
		episodes, err := episodeRepository.GetAllUpcomingEpisodes(series)
		if err != nil {
			InternalServerErrorHandler(response, request, err)
			return
		}
		if len(episodes) > 0 {
			series.UnwatchedEpisodes = episodes
			resultSlice = append(resultSlice, series)
		}
	}

	json.NewEncoder(response).Encode(resultSlice)
	logAccess(database, request)
}

/*
{
	"ID": 45,
	"Title": "Parks and Recreation",
	"ProviderURL": "https://www.themoviedb.org/tv/8592-parks-and-recreation",
	"current_episode": {
		"ID": 3821,
		"Episode": 3,
		"Season": 6,
		"Title": "The Pawnee-Eagleton Tip off Classic",
		"Description": "Ben, Leslie and Chris meet with Eagleton city councilor Ingrid de Forest to discuss financial matters; Ann brings April to vet school orientation; Ron tries to get off the grid.",
		"ReleaseDate": "2013-10-03T00:00:00Z"
	}
}
 */
func MovePointer(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	seriesName := vars["series"]
	request.ParseForm()
	episodeIdSlice, passed := request.Form["episode_id"]
	if !passed {
		BadRequestHandler(response, request)
		return
	}
	targetEpisodeId := episodeIdSlice[0]
	_, database := getSettingsAndDatabase(response, request)
	defer database.Close()
	seriesRepository := Models.SeriesRepository{Db: database}
	episodeRepository := Models.EpisodeRepository{Db: database}

	series, err := seriesRepository.GetByName(seriesName, true)
	if err != nil {
		if err == sql.ErrNoRows {
			NotFoundHandler(response, request)
		} else {
			InternalServerErrorHandler(response, request, err)
		}
		return
	}

	id, err := strconv.ParseInt(targetEpisodeId, 10, 64)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
	}

	episode, err := episodeRepository.GetById(id, true)
	if err != nil {
		NotFoundErrorHandler(response, request, errors.New("No episode with that id found"))
		return
	}
	series.WatchPointer = episode
	err = seriesRepository.UpdateWatchPointer(series)
	if err != nil {
		InternalServerErrorHandler(response, request, err)
		return
	}

	json.NewEncoder(response).Encode(series)
	logAccess(database, request)
}
