package Models

import (
	"database/sql"
	"time"
	"fmt"
	"os"
	"github.com/fatih/color"
)

const (
	LogTableName        = "Log"
	LogBasicSelectQuery = "SELECT `Type`, `Message`, `Time`, `Caller`, `RequestUri`, `StatusCode` FROM " + LogTableName
	LogTypeMessage      = 1
	LogTypeWarning      = 2
	LogTypeError        = 3
)

type Log struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	Time       time.Time `json:"time"`
	Type       int       `json:"type"`
	Caller     string    `json:"caller"`
	RequestUri string    `json:"request_uri"`
	StatusCode int       `json:"status_code"`
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

	currentTime := time.Now()
	printMessage := "[" + currentTime.Format("2006-01-02 15:04:05") + "] "

	switch log.Type {
	case LogTypeMessage:
		fmt.Fprintln(
			os.Stdout,
			color.GreenString(printMessage),
			color.GreenString(log.Caller),
			color.GreenString("[Info]"),
			color.GreenString(log.Message),
			color.GreenString(log.RequestUri),
		)
		break
	case LogTypeWarning:
		fmt.Fprintln(
			os.Stdout,
			color.YellowString(printMessage),
			color.YellowString(log.Caller),
			color.YellowString("[Warning]"),
			color.YellowString(log.Message),
			color.YellowString(log.RequestUri),
		)
		break
	case LogTypeError:
		fmt.Fprintln(
			os.Stderr,
			color.RedString(printMessage),
			color.RedString(log.Caller),
			color.RedString("[Error]"),
			color.RedString(log.Message),
			color.RedString(log.RequestUri),
		)
		break
	default:
		fmt.Fprintln(
			os.Stdout,
			printMessage,
			"[Unknown]",
			log.Message,
		)
	}

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
