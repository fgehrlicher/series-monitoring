package Server

import (
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Config"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/fatih/color"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Endpoints"
	"time"
	"os"
	"os/signal"
	"context"
	"github.com/gorilla/handlers"
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func Init() {
	settings, err := Config.GetConfiguration()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	router := mux.NewRouter()

	router.StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(Endpoints.NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(Endpoints.MethodNotAllowedHandler)

	/*
		root endpoint.

		HTTP METHOD:
		GET

		Url:
		/

		Example Requests:
		http://127.0.0.1:5000

		Returns:
		{
			"message": "Series watcher v3"
		}
	*/
	router.HandleFunc("/", Endpoints.RootEndpoint).Methods("GET")

	/*
		returns all log entries.

		HTTP METHOD:
		GET

		Url:
		/logs/

		Url Parameter:
		"since", int, max age of the log entries in seconds

		Example Requests:
		http://127.0.0.1:5000/logs/
		http://127.0.0.1:5000/logs/?since=1000

		Returns:
		[
			{
				"message": "Authorized Call",
				"time": "2018-03-12T07:24:21Z",
				"type": 1,
				"caller": "127.0.0.1:53419",
				"request_uri": "\/logs\/",
				"status_code": 200,
				"id": 0
			},
			....
		]
 	*/
	router.HandleFunc("/logs/", Endpoints.GetAllLogs).Methods("GET")

	/*
		returns all message log entries.

		HTTP METHOD:
		GET

		Url:
		/logs/message/

		Url Parameter:
		"since", int, max age of the log entries in seconds

		Example Requests:
		http://127.0.0.1:5000/logs/message/
		http://127.0.0.1:5000/logs/message/?since=1000

		Returns:
		[
			{
				"message": "Authorized Call",
				"time": "2018-03-12T07:24:21Z",
				"type": 1,
				"caller": "127.0.0.1:53419",
				"request_uri": "\/logs\/",
				"status_code": 200,
				"id": 0
			},
			....
		]
 	*/
	router.HandleFunc("/logs/message/", Endpoints.GetAllMessages).Methods("GET")

	/*
		returns all warning log entries.

		HTTP METHOD:
		GET

		Url:
		/logs/warning/

		Url Parameter:
		"since", int, max age of the log entries in seconds

		Example Requests:
		http://127.0.0.1:5000/logs/message/
		http://127.0.0.1:5000/logs/message/?since=1000

		Returns:
		[
			{
				"message": "Not Found",
				"time": "2018-03-09T13:51:01Z",
				"type": 2,
				"caller": "127.0.0.1:58052",
				"request_uri": "\/logs\/adasd",
				"status_code": 404,
				"id": 0
			},
			....
		]
 	*/
	router.HandleFunc("/logs/warning/", Endpoints.GetAllWarnings).Methods("GET")

	/*
		returns all error log entries.

		HTTP METHOD:
		GET

		Url:
		/logs/error/

		Url Parameter:
		"since", int, max age of the log entries in seconds

		Example Requests:
		http://127.0.0.1:5000/logs/error/
		http://127.0.0.1:5000/logs/error/?since=1000

		Returns:
		[
			{
				"message": "Server Error",
				"time": "2018-03-09T13:51:01Z",
				"type": 3,
				"caller": "127.0.0.1:58052",
				"request_uri": "\/logs\/",
				"status_code": 500,
				"id": 0
			},
			....
		]
 	*/
	router.HandleFunc("/logs/error/", Endpoints.GetAllErrors).Methods("GET")

	/*
		returns all providers.

		HTTP METHOD:
		GET

		Url:
		/provider/

		Example Requests:
		http://127.0.0.1:5000/provider/

		Returns:
		[
			{
				"id": 1,
				"name": "Netflix",
				"image_path": "provider\/netflix\/image"
			},
			...
		]
 	*/
	router.HandleFunc("/provider/", Endpoints.GetAllProviders).Methods("GET")

	/*
		returns one provider.

		HTTP METHOD:
		GET

		Url:
		/provider/{provider}/

		Path Parameter:
		"provider", string, provider name

		Example Requests:
		http://127.0.0.1:5000/provider/netflix/

		Returns:
		{
			"id": 1,
			"name": "Netflix",
			"image_path": "provider\/netflix\/image"
		}
 	*/
	router.HandleFunc("/provider/{provider}/", Endpoints.GetProvider).Methods("GET")

	/*
		returns the provider image.

		HTTP METHOD:
		GET

		Url:
		/provider/{provider}/image/

		Path Parameter:
		"provider", string, provider name

		Example Requests:
		http://127.0.0.1:5000/provider/netflix/image/

		Returns:
		Content-Type: image/png
	*/
	router.HandleFunc("/provider/{provider}/image/", Endpoints.GetProviderImage).Methods("GET")

	/*
		returns all series.

		HTTP METHOD:
		GET

		Url:
		/series/

		Example Requests:
		http://127.0.0.1:5000/series/

		Returns:
		[
			{
				"id": 10,
				"Title": "Marvel's Daredevil",
				"image_path": "series/Marvel%27s%20Daredevil/image/"
				"ProviderURL": "https://www.themoviedb.org/tv/61889-daredevil",
				"UnwatchedEpisodes": null,
				"current_episode": {
					"id": 643,
					"Episode": 13,
					"Season": 2,
					"Title": "A Cold Day in Hell's Kitchen",
					"Description": "In the season finale, Daredevil is backed into the ultimate showdown for his own life -- and the future of Hell's Kitchen.",
					"release_date": "2016-03-18T00:00:00Z",
					"ImageUrl": "/series/Marvel%27s%20Daredevil/season/2/episode/13/image/"
				},
			},
			....
		]
	*/
	router.HandleFunc("/series/", Endpoints.GetAllSeries).Methods("GET")

	/*
		creates a new series.

		HTTP METHOD:
		POST

		Url:
		/series/

		Form Parameter:
		"series_url", string, tmdb url

		Example Requests:
		http://127.0.0.1:5000/series/
		x-www-form-urlencoded
		key: series_url value: https://www.themoviedb.org/tv/74204-big-mouth

		Returns:
		{
			"id": 57,
			"Title": "Big Mouth",
			"image_path": "series/Big%20Mouth/image/",
			"ProviderURL": "https://www.themoviedb.org/tv/74204-big-mouth",
			"UnwatchedEpisodes": null,
			"current_episode": {
				"id": 4511,
				"Episode": 1,
				"Season": 1,
				"Title": "Ejaculation",
				"Description": "As Andrew falls under the spell of the randy Hormone Monster, his buddy Nick becomes obsessed with the lack of changes in his own body.",
				"release_date": "2017-09-29T00:00:00Z",
				"ImageUrl": "/series/Big%20Mouth/season/1/episode/1/image/"
			}
		}
	*/
	router.HandleFunc("/series/", Endpoints.CreateSeries).Methods("POST")

	/*
		returns all umwatched episodes.

		HTTP METHOD:
		GET

		Url:
		/series/*\/unwatched/

		Example Requests:
		http://127.0.0.1:5000/series/*\/unwatched/

		Returns:
		[
			{
				id: 8,
				Title: "Archer",
				image_path: "series/Archer/image/",
				ProviderURL: "https://www.themoviedb.org/tv/10283-archer",
				UnwatchedEpisodes: [
					{
						id: 4371,
						Episode: 4,
						Season: 9,
						Title: "A Warrior in Costume",
						Description: "Archer keeps an eye out for a past nemesis.",
						release_date: "2018-05-16T00:00:00Z",
						ImageUrl: "/series/Archer/season/9/episode/4/image/"
					},
					....
				],
				current_episode: {
					id: 3863,
					Episode: 3,
					Season: 9,
					Title: "Archer: Danger Island – Different Modes of Preparing the Fruit",
					Description: "Archer, Pam, and Crackers try to find the missing key to their latest get-rich-quick scheme.",
					release_date: "2018-05-09T00:00:00Z",
					ImageUrl: "/series/Archer/season/9/episode/3/image/"
				}
			},
			....
		]
	*/
	router.HandleFunc("/series/*/unwatched/", Endpoints.GetAllNewEpisodes).Methods("GET")

	/*
		returns all upcoming episodes.

		HTTP METHOD:
		GET

		Url:
		/series/*\/upcoming/

		Example Requests:
		http://127.0.0.1:5000/series/*\/upcoming/

		Returns:
		[
			{
				id: 8,
				Title: "Archer",
				image_path: "series/Archer/image/",
				ProviderURL: "https://www.themoviedb.org/tv/10283-archer",
				UnwatchedEpisodes: [
					{
						id: 4371,
						Episode: 4,
						Season: 9,
						Title: "A Warrior in Costume",
						Description: "Archer keeps an eye out for a past nemesis.",
						release_date: "2018-05-16T00:00:00Z",
						ImageUrl: "/series/Archer/season/9/episode/4/image/"
					},
					....
				],
				current_episode: {
					id: 3863,
					Episode: 3,
					Season: 9,
					Title: "Archer: Danger Island – Different Modes of Preparing the Fruit",
					Description: "Archer, Pam, and Crackers try to find the missing key to their latest get-rich-quick scheme.",
					release_date: "2018-05-09T00:00:00Z",
					ImageUrl: "/series/Archer/season/9/episode/3/image/"
				}
			},
			....
		]
	*/
	router.HandleFunc("/series/*/upcoming/", Endpoints.GetAllUpcomingEpisodes).Methods("GET")

	/*
		updates all series.

		HTTP METHOD:
		POST

		Url:
		/series/*\/update/

		Example Requests:
		http://127.0.0.1:5000/series/*\/update/

		Returns:
		{
			"message": "Success",
			"Steps": [
				{
					"message": "Imported 4 Episodes for Series: Archer"
				},
				....
			]
		}
	*/
	router.HandleFunc("/series/*/update/", Endpoints.UpdateAllSeries).Methods("POST")

	/*
	@TODO
	 */
	router.HandleFunc("/series/*/integrity-check/", Endpoints.CheckIntegrity).Methods("POST")

	/*
		returns one series.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/

		Path Parameter:
		"series", string, series name

		Example Requests:
		http://127.0.0.1:5000/series/archer/

		Returns:
		{
			id: 8,
			Title: "Archer",
			image_path: "series/Archer/image/",
			ProviderURL: "https://www.themoviedb.org/tv/10283-archer",
			UnwatchedEpisodes: null,
			current_episode: {
				id: 3863,
				Episode: 3,
				Season: 9,
				Title: "Archer: Danger Island – Different Modes of Preparing the Fruit",
				Description: "Archer, Pam, and Crackers try to find the missing key to their latest get-rich-quick scheme.",
				release_date: "2018-05-09T00:00:00Z",
				ImageUrl: "/series/Archer/season/9/episode/3/image/"
			}
		}
	*/
	router.HandleFunc("/series/{series}/", Endpoints.GetSeries).Methods("GET")

	/*
		updates a series pointer for one series.

		HTTP METHOD:
		POST

		Url:
		/series/{series}/pointer/

		Path Parameter:
		"series", string, series name

		Form Parameter:
		"episode_id", int, the episode id

		Example Requests:
		http://127.0.0.1:5000/series/Big%20Mouth/pointer/
		x-www-form-urlencoded
		key: episode_id value: 1

		Returns:
		{
			"id": 57,
			"Title": "Big Mouth",
			"image_path": "series/Big%20Mouth/image/",
			"ProviderURL": "https://www.themoviedb.org/tv/74204-big-mouth",
			"UnwatchedEpisodes": null,
			"current_episode": {
				"id": 4511,
				"Episode": 1,
				"Season": 1,
				"Title": "Ejaculation",
				"Description": "As Andrew falls under the spell of the randy Hormone Monster, his buddy Nick becomes obsessed with the lack of changes in his own body.",
				"release_date": "2017-09-29T00:00:00Z",
				"ImageUrl": "/series/Big%20Mouth/season/1/episode/1/image/"
			}
		}
	*/
	router.HandleFunc("/series/{series}/pointer/", Endpoints.MovePointer).Methods("POST")

	/*
		updates one series.

		HTTP METHOD:
		POST

		Url:
		/series/{series}/update/

		Path Parameter:
		"series", string, series name

		Example Requests:
		http://127.0.0.1:5000/series/archer/update/

		Returns:
		{
			"message": "Success",
			"Steps": [
				{
					"message": "Imported 4 Episodes for Series: Archer"
				}
			]
		}
	*/
	router.HandleFunc("/series/{series}/update/", Endpoints.UpdateSeries).Methods("POST")

	/*
		returns the series image.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/image/

		Path Parameter:
		"series", string, series name

		Example Requests:
		http://127.0.0.1:5000/series/archer/image/

		Returns:
		Content-Type: image/png
	*/
	router.HandleFunc("/series/{series}/image/", Endpoints.GetSeriesImage).Methods("GET")

	/*
		returns all episodes for one series.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/season/*\/episodes/

		Path Parameter:
		"series", string, series name

		Example Requests:
		http://127.0.0.1:5000/series/archer/season/*\/episodes/

		Returns:
		[
			{
				id: 399,
				Episode: 1,
				Season: 1,
				Title: "Mole Hunt",
				Description: "Archer is in trouble with his Mother and the Comptroller because his expense account is way out of proportion to his actual expenses.",
				release_date: "2009-09-16T00:00:00Z",
				ImageUrl: "/series/Archer/season/1/episode/1/image/"
			},
			...
		]
	*/
	router.HandleFunc("/series/{series}/season/*/episodes/", Endpoints.GetAllEpisodes).Methods("GET")

	/*
		returns all unwatched episodes for one series.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/season/*\/episodes/unwatched/

		Path Parameter:
		"series", string, series name

		Example Requests:
		http://127.0.0.1:5000/series/archer/season/*\/episodes/unwatched/

		Returns:
		[
			{
				id: 399,
				Episode: 1,
				Season: 1,
				Title: "Mole Hunt",
				Description: "Archer is in trouble with his Mother and the Comptroller because his expense account is way out of proportion to his actual expenses. So he creates the idea that a Mole has breached ISIS and he needs to get into the mainframe to hunt him down (so he can cover his tracks!). All this leads to a surprising ending.",
				release_date: "2009-09-16T00:00:00Z",
				ImageUrl: "/series/Archer/season/1/episode/1/image/"
			},
			...
		]
	*/
	router.HandleFunc("/series/{series}/season/*/episodes/unwatched/", Endpoints.GetNewEpisodes).Methods("GET")

	/*
		returns all upcoming episodes for one series.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/season/*\/episodes/upcoming/

		Path Parameter:
		"series", string, series name

		Example Requests:
		http://127.0.0.1:5000//series/archer/season/*\/episodes/upcoming/

		Returns:
		[
			{
				id: 399,
				Episode: 1,
				Season: 1,
				Title: "Mole Hunt",
				Description: "Archer is in trouble with his Mother and the Comptroller because his expense account is way out of proportion to his actual expenses. So he creates the idea that a Mole has breached ISIS and he needs to get into the mainframe to hunt him down (so he can cover his tracks!). All this leads to a surprising ending.",
				release_date: "2009-09-16T00:00:00Z",
				ImageUrl: "/series/Archer/season/1/episode/1/image/"
			},
			...
		]
	*/
	router.HandleFunc("/series/{series}/season/*/episodes/upcoming/", Endpoints.GetUpcomingEpisodes).Methods("GET")

	/*
		returns one season for one series.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/season/{season}/

		Path Parameter:
		"series", string, series name
		"season", int, season number

		Example Requests:
		http://127.0.0.1:5000/series/archer/season/1/

		Returns:
		[
			{
				id: 399,
				Episode: 1,
				Season: 1,
				Title: "Mole Hunt",
				Description: "Archer is in trouble with his Mother and the Comptroller because his expense account is way out of proportion to his actual expenses. So he creates the idea that a Mole has breached ISIS and he needs to get into the mainframe to hunt him down (so he can cover his tracks!). All this leads to a surprising ending.",
				release_date: "2009-09-16T00:00:00Z",
				ImageUrl: "/series/Archer/season/1/episode/1/image/"
			},
			...
		]
	*/
	router.HandleFunc("/series/{series}/season/{season}/", Endpoints.GetAllEpisodesBySeason).Methods("GET")

	/*
		returns one episode.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/season/{season}/episode/{episode}/

		Path Parameter:
		"series", string, series name
		"season", int, season number
		"episode", int, episode number

		Example Requests:
		http://127.0.0.1:5000/series/archer/season/1/episode/1/

		Returns:
		{
			id: 399,
			Episode: 1,
			Season: 1,
			Title: "Mole Hunt",
			Description: "Archer is in trouble with his Mother and the Comptroller because his expense account is way out of proportion to his actual expenses. So he creates the idea that a Mole has breached ISIS and he needs to get into the mainframe to hunt him down (so he can cover his tracks!). All this leads to a surprising ending.",
			release_date: "2009-09-16T00:00:00Z",
			ImageUrl: "/series/Archer/season/1/episode/1/image/"
		}
	*/
	router.HandleFunc("/series/{series}/season/{season}/episode/{episode}/", Endpoints.GetEpisode).Methods("GET")

	/*
		returns the episode image.

		HTTP METHOD:
		GET

		Url:
		/series/{series}/season/{season}/episode/{episode}/image/

		Path Parameter:
		"series", string, series name
		"season", int, season number
		"episode", int, episode number

		Example Requests:
		http://127.0.0.1:5000/series/archer/season/1/episode/1/image/

		Returns:
		Content-Type: image/png
	*/
	router.HandleFunc("/series/{series}/season/{season}/episode/{episode}/image/", Endpoints.GetEpisodeImage).Methods("GET")

	corsOptions := handlers.AllowedOrigins([]string{"*"})
	ip := settings.ServerSettings.Ip
	port := settings.ServerSettings.Port
	server := &http.Server{
		Addr:         ip + ":" + port,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 360,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(corsOptions)(router),
	}

	figure.NewFigure(" SERIES WATCHER ", "", true).Print()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			server.ListenAndServe()
		}
	}()

	color.Green("\nServer listening on: " + ip + ":" + port + "\n\n")

	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt)

	<-osSignalChannel

	deadLineContext, cancelFunction := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunction()
	server.Shutdown(deadLineContext)
	fmt.Print("Shutting down... \n")
	os.Exit(0)

}
