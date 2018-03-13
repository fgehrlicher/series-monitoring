package Parser

const (
	TMDbBaseUrl       = "https://www.themoviedb.org"
	CoverImagesPath   = "Series/"
	EpisodeImagesPath = "Series/"
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
	init() error
	GetSeriesTitle() (string, error)
	GetSeriesCover() (string, error)
}
