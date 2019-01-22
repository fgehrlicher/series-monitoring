package Models

import "database/sql"

type Url struct {
	EpisodeID  int
	ProviderID int
	URL        sql.NullString
	ID         int
}
