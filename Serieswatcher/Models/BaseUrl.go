package Models

import "database/sql"

type BaseUrl struct {
	ProviderID int
	SeriesID   int
	URL        sql.NullString
	ID         int
}
