package Maintenance

import "database/sql"

type SynchronizeController struct {
	Db *sql.DB
}

func NewSynchronizeController(db *sql.DB) *SynchronizeController {
	return &SynchronizeController{Db: db}
}

func (Controller *SynchronizeController) run() error {


	return nil
}
