package Models

import (
	"database/sql"
	"time"
)

type ManualAction struct {
	Message sql.NullString
	Time    time.Time
	Type    sql.NullString
	Done    sql.NullInt64
	ID      int
}
