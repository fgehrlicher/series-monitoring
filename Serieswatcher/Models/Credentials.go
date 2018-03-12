package Models

import "database/sql"

type Credentials struct {
	Password sql.NullString
	ID       int
}
