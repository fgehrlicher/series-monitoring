package Models

import (
	"database/sql"
	"gitea.fge.cloud/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Config"
	"errors"
)

func GetDatabaseConnection(settings Config.Settings) (*sql.DB, error)  {

	dataSourceName := settings.DatabaseSettings.User + ":" +
		settings.DatabaseSettings.Password + "@" +
		"tcp(" + settings.DatabaseSettings.Host + ":" + settings.DatabaseSettings.Port + ")/" +
		settings.DatabaseSettings.Database + "?parseTime=true"
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, errors.New("db error")
	}

	return db, nil

}
