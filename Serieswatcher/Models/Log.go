package Models

import (
	"database/sql"
	"time"
)

const (
	LogTableName        = "Log"
	LogBasicSelectQuery = "SELECT `Type`, `Message`, `Time`, `Caller`, `RequestUri`, `StatusCode` FROM " + LogTableName
	LogTypeMessage      = 1
	LogTypeWarning      = 2
	LogTypeError        = 3
)

type Log struct {
	Message    string
	Time       time.Time
	Type       int
	Caller     string
	RequestUri string
	StatusCode int
	ID         int
}

type LogRepository struct {
	Db *sql.DB
}

func (repository *LogRepository) Scan(row *sql.Rows) Log {
	log := Log{}
	row.Scan(&log.Type, &log.Message, &log.Time, &log.Caller, &log.RequestUri, &log.StatusCode)
	return log
}

func (repository *LogRepository) Persist(log Log) error {
	query := "INSERT INTO " + LogTableName + " (Message, Time, Type, Caller, RequestUri, StatusCode) VALUES (?, NOW(), ?, ?, ?, ?)"
	_, err := repository.Db.Exec(
		query,
		log.Message,
		log.Type,
		log.Caller,
		log.RequestUri,
		log.StatusCode,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repository *LogRepository) GetAll() ([]Log, error) {
	var logs []Log
	row, err := repository.Db.Query(LogBasicSelectQuery)
	if err != nil {
		return logs, err
	}
	defer row.Close()
	for row.Next() {
		logs = append(logs, repository.Scan(row))
	}
	return logs, err
}

func (repository *LogRepository) GetAllByType(logType int) ([]Log, error) {
	var logs []Log
	row, err := repository.Db.Query(LogBasicSelectQuery+" WHERE `Type` = ?", logType)
	if err != nil {
		return logs, err
	}
	defer row.Close()
	for row.Next() {
		logs = append(logs, repository.Scan(row))
	}
	return logs, err
}

func (repository *LogRepository) GetAllSince(since int) ([]Log, error) {
	var logs []Log
	row, err := repository.Db.Query(
		LogBasicSelectQuery+" WHERE `Time` >= DATE_SUB(current_timestamp, INTERVAL ? SECOND)",
		since,
	)
	if err != nil {
		return logs, err
	}
	defer row.Close()
	for row.Next() {
		logs = append(logs, repository.Scan(row))
	}
	return logs, err
}

func (repository *LogRepository) GetAllSinceByType(since int, logType int) ([]Log, error) {
	var logs []Log
	row, err := repository.Db.Query(
		LogBasicSelectQuery+" WHERE `Time` >= DATE_SUB(current_timestamp, INTERVAL ? SECOND) AND `Type` = ?",
		since,
		logType,
	)
	defer row.Close()
	if err != nil {
		return logs, err
	}
	for row.Next() {
		logs = append(logs, repository.Scan(row))
	}
	return logs, err
}
