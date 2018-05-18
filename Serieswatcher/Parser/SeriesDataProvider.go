package Parser

import "bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Models"

const (
	CoverImagesPath   = "Series/"
	EpisodeImagesPath = "Episode/"
)

type ParserError struct {
	text string
}

func (error *ParserError) Error() string {
	return error.text
}

func NewError(text string) *ParserError {
	return &ParserError{text: "parser error: " + text}
}

type SeriesDataProvider interface {
	GetSeries() (*Models.Series, error)
	GetAllNewEpisodes(Models.Episode) ([]Models.Episode, error)
	GetEpisode(season int, episode int) (*Models.Episode, error, error)
}

type SeriesContentProvider interface {
}
