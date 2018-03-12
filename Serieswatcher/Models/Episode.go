package Models

import (
	"database/sql"
	"time"
)

type Episode struct {
	Description sql.NullString
	Episode     sql.NullInt64
	ImageID     int
	ReleaseDate time.Time
	Season      sql.NullInt64
	SeriesID    int
	Title       sql.NullString
	ID          int
}
