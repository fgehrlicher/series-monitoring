package Models

import "database/sql"

type Series struct {
	ImageID      int
	ProviderType sql.NullString
	ProviderURL  sql.NullString
	Title        sql.NullString
	ID           int
}
