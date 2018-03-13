package Parser

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"path"
	"net/http"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Config"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"
)

type TMDbHandler struct {
	BaseUrl     string
	SeriesUrl   string
	SeriesTitle string
	CoverImage  Models.Image
	Document    *goquery.Document
}

func NewTMDbHandler(seriesUrl string, settings Config.Settings) (*TMDbHandler, error) {
	handler := &TMDbHandler{
		SeriesUrl:  seriesUrl,
		CoverImage: Models.Image{Settings: settings},
	}
	err := handler.init()
	return handler, err
}

func (handler *TMDbHandler) init() error {
	if !strings.HasPrefix(handler.SeriesUrl, TMDbBaseUrl) {
		return NewError("Invalid Series Url passed")
	}
	doc, err := goquery.NewDocument(handler.SeriesUrl)
	if err != nil {
		return err
	}
	handler.Document = doc
	return nil
}

func (handler *TMDbHandler) GetSeries() (Models.Series, error) {

	series := Models.Series{}

	title, err := handler.GetSeriesTitle()
	if err != nil {
		return series, err
	}
	image, err := handler.GetSeriesCover()
	if err != nil {
		return series, err
	}

	series.Title = title
	series.ImagePath = image.RelativePath
	series.ImageOriginUrl = image.OriginURL
	series.ProviderURL = handler.SeriesUrl

	return series, nil
}

func (handler *TMDbHandler) GetSeriesTitle() (string, error) {
	titleNode := handler.Document.Find("div.header_poster_wrapper section div.title span a h2")
	title := titleNode.Text()
	if title == "" {
		return "", NewError("header node not found")
	}
	handler.SeriesTitle = title
	return handler.SeriesTitle, nil
}

func (handler *TMDbHandler) GetSeriesCover() (*Models.Image, error) {
	if handler.SeriesTitle == "" {
		handler.GetSeriesTitle()
	}

	seriesTitle := handler.SeriesTitle
	fileName := strings.ToLower(strings.Replace(seriesTitle, " ", "-", -1))

	relativeCoverImagePath := path.Join(CoverImagesPath, fileName)
	handler.CoverImage.RelativePath = relativeCoverImagePath

	if handler.CoverImage.Exists() {
		handler.CoverImage.Delete()
	}

	coverImageNode := handler.Document.Find("div.poster div.image_content img.poster")
	coverImageUrl, exists := coverImageNode.Attr("src")
	if !exists {
		return &handler.CoverImage, NewError("cover image link not found")
	}

	response, err := http.Get(coverImageUrl)
	if err != nil {
		return &handler.CoverImage, err
	}

	handler.CoverImage.Create(response.Body)
	handler.CoverImage.OriginURL = coverImageUrl
	defer response.Body.Close()

	return &handler.CoverImage, nil
}
