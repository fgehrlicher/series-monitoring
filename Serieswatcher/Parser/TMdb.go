package Parser

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"path"
	"net/http"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
	"strconv"
	"github.com/murlokswarm/errors"
	"time"
)

const (
	TMDbBaseUrl               = "https://www.themoviedb.org"
	TMDbDefaultParameter      = "?language=en-US"
	TMDbPageNotFound          = "Oops!—We can't find the page you're looking for."
	TMDbNoImageOnBackdropPage = "There no backdrop stills added to this entry."
	TMDbNoImagesOnOverview    = "No episode images have been added."
	TMDbNoDescription         = "We don't have an overview translated in English. Help us expand our database by adding one."
)

type TMDbHandler struct {
	SeriesUrl string
	Series    *Models.Series
	Document  *goquery.Document
}

func NewTMDbHandler(series Models.Series) (*TMDbHandler, error) {
	handler := &TMDbHandler{
		SeriesUrl: series.ProviderURL,
		Series:    &series,
	}
	handler.Series.Image = &Models.Image{ImageType: Models.ImageSeries}
	if !strings.HasPrefix(handler.SeriesUrl, TMDbBaseUrl) {
		return nil, NewError("Invalid Series Url passed")
	}
	doc, err := goquery.NewDocument(handler.SeriesUrl + TMDbDefaultParameter)
	if err != nil {
		return nil, err
	}
	handler.Series.ProviderURL = handler.SeriesUrl
	handler.Document = doc
	return handler, err
}

func (handler *TMDbHandler) GetAllNewEpisodes(lastEpisode Models.Episode) ([]Models.Episode, error) {
	var episodes []Models.Episode

	season := lastEpisode.Season
	episode := lastEpisode.Episode + 1

	err := handler.navigateToSeason(season)
	if err != nil {
		return episodes, err
	}
	for {
		err := handler.navigateToSeason(season)
		if handler.isErrorPage(handler.Document) || err != nil {
			break
		}
		for {
			episodeStruct, episodeNotFoundErr, err := handler.GetEpisode(season, episode)
			if episodeNotFoundErr != nil {
				break
			}
			if err != nil {
				return episodes, err
			}
			episodes = append(episodes, *episodeStruct)
			episode ++

		}
		episode = 1
		season ++
	}

	return episodes, nil
}

func (handler *TMDbHandler) GetEpisode(season int, episode int) (*Models.Episode, error, error) {
	baseUrl := handler.SeriesUrl +
		"/season/" + strconv.Itoa(season) +
		"/episode/" + strconv.Itoa(episode)
	episodeUrl := baseUrl + TMDbDefaultParameter
	imageDetailUrl := baseUrl + "/images/backdrops" + TMDbDefaultParameter
	episodeElement := &Models.Episode{
		Image: &Models.Image{
			ImageType: Models.ImageEpisode,
		},
		Episode: episode,
		Season:  season,
	}
	episodePage, err := goquery.NewDocument(episodeUrl)
	if err != nil {
		return nil, nil, err
	}
	if handler.isErrorPage(episodePage) {
		return nil, NewError("Episode Not Found"), nil
	}
	episodeSelection := episodePage.Find("div.opened")
	if len(episodeSelection.Nodes) != 1 {
		return nil, NewError("Invalid Episode"), nil
	}

	titleNode := episodeSelection.Find("div.info div div.title div.wrapper h3 a.open")
	title := titleNode.Text()
	episodeElement.Title = title
	descriptionNode := episodeSelection.Find("div.info div div.overview p")
	description := descriptionNode.Text()
	if description == TMDbNoDescription {
		return nil, NewError("No Description Jet"), nil
	}
	episodeElement.Description = description

	dateNode := episodeSelection.Find("div.info div.title div.date")
	releaseDate := dateNode.Text()
	if releaseDate != "" {
		episodeElement.ReleaseDate, err = time.Parse("January 2, 2006", releaseDate)
		if err != nil {
			return nil, nil, err
		}
	}

	imageDetailPage, err := goquery.NewDocument(imageDetailUrl)
	if err != nil {
		return nil, nil, err
	}
	if handler.isErrorPage(imageDetailPage) {
		return nil, nil, NewError("Image Not Found")
	}

	var (
		imageUrl string
		exists   bool
	)
	if handler.isEmptyImagePage(imageDetailPage) {
		imageNode := episodeSelection.Find("div.image a.open img")
		imageUrl, exists = imageNode.Attr("data-src")
		if !exists {
			if handler.hasImages(episodeSelection) {
				return nil, nil, NewError("Image Not Found, Url: " + episodeUrl)
			}
		}
	} else {
		imageNode := imageDetailPage.Find("div.results ul li div.image_content a.image")
		imageUrl, exists = imageNode.Attr("href")
		if !exists {
			return nil, nil, NewError("Image Not Found, Url: " + imageDetailUrl)
		}
	}

	if imageUrl == "" {
		return episodeElement, nil, nil
	}
	response, err := http.Get(imageUrl)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()
	episodeElement.Image.OriginURL = imageUrl
	fileName := strings.ToLower(strings.Replace(handler.Series.Title, " ", "-", -1)) +
		strconv.Itoa(season) + "-" + strconv.Itoa(episode)
	relativeCoverImagePath := path.Join(EpisodeImagesPath, fileName)
	episodeElement.Image.RelativePath = relativeCoverImagePath
	err = episodeElement.Image.CreateFile(response.Body)
	if err != nil {
		return nil, nil, err
	}
	episodeElement.Image.OriginURL = imageUrl
	return episodeElement, nil, nil
}

func (handler *TMDbHandler) GetSeries() (*Models.Series, error) {
	var err error
	err = handler.getSeriesTitle()
	if err != nil {
		return nil, err
	}
	err = handler.getSeriesCover()
	if err != nil {
		return nil, err
	}
	episode, notFoundError, err := handler.GetEpisode(1, 1)
	if err != nil {
		return nil, err
	}
	if notFoundError != nil {
		return nil, notFoundError
	}
	handler.Series.WatchPointer = episode
	return handler.Series, nil
}

func (handler *TMDbHandler) getSeriesTitle() error {
	titleNode := handler.Document.Find("div.header_poster_wrapper section div.title span a h2")
	title := titleNode.Text()
	if title == "" {
		return NewError("header node not found")
	}
	handler.Series.Title = title
	return nil
}

func (handler *TMDbHandler) getSeriesCover() error {
	coverImageNode := handler.Document.Find("div.poster div.image_content img.poster")
	coverImageUrl, exists := coverImageNode.Attr("src")
	if !exists {
		return NewError("cover image link not found")
	}

	fileName := strings.ToLower(strings.Replace(handler.Series.Title, " ", "-", -1))
	relativeCoverImagePath := path.Join(CoverImagesPath, fileName)
	handler.Series.Image.RelativePath = relativeCoverImagePath

	httpClient := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	response, err := httpClient.Get(coverImageUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	handler.Series.Image.CreateFile(response.Body)
	handler.Series.Image.OriginURL = coverImageUrl

	return nil
}

func (handler *TMDbHandler) isErrorPage(document *goquery.Document) bool {
	isErrorPage := false
	document.Find("h2").Each(func(_ int, element *goquery.Selection) {
		if element.Text() == TMDbPageNotFound {
			isErrorPage = true
		}
	})
	return isErrorPage
}

func (handler *TMDbHandler) hasImages(selection *goquery.Selection) bool {
	hasImages := true
	elements := selection.Find("div.expanded_info.wrap p")
	elements.Each(func(_ int, element *goquery.Selection) {
		text := element.Text()
		if text == TMDbNoImagesOnOverview {
			hasImages = false
		}
	})
	return hasImages
}

func (handler *TMDbHandler) isEmptyImagePage(document *goquery.Document) bool {
	isErrorPage := false
	document.Find("p").Each(func(_ int, element *goquery.Selection) {
		if element.Text() == TMDbNoImageOnBackdropPage {
			isErrorPage = true
		}
	})
	return isErrorPage
}

func (handler *TMDbHandler) navigateToSeason(season int) error {
	doc, err := goquery.NewDocument(handler.SeriesUrl + "/season/" + strconv.Itoa(season) + TMDbDefaultParameter)
	if err != nil {
		return err
	}
	if handler.isErrorPage(doc) {
		return errors.New("Season Not Found")
	}
	handler.Document = doc
	return nil
}
