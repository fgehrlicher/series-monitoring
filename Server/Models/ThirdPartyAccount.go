package Models

import "database/sql"

type ThirdPartyAccount struct {
	CredentialsID int
	ProviderID    int
	UserID        int
	Username      sql.NullString
	ID            int
}
