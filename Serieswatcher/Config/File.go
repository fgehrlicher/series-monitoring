package Config

import (
	"os"
	"encoding/json"
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net"
)

type DatabaseSettings struct {
	Host      string
	Container string `json:"container-alias"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
}

type ServerSettings struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type Settings struct {
	DatabaseSettings DatabaseSettings `json:"database"`
	ServerSettings   ServerSettings   `json:"server"`
}

func GetConfiguration() (Settings, error) {
	var settings Settings
	configFile, err := os.Open("./config.json")
	defer configFile.Close()
	if err != nil {
		return settings, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&settings)
	if err != nil {
		return settings, err
	}
	err = validate(&settings)
	if err != nil {
		return settings, err
	}
	err = testConnection(settings)
	if err != nil {
		return settings, err
	}
	return settings, nil
}

func validate(settings *Settings) error {

	if settings.DatabaseSettings.Container == "" {
		return errors.New("No Container Alias specified ")
	}
	ip, _ := net.LookupIP(settings.DatabaseSettings.Container)
	if ip == nil {
		return errors.New("Database Container not found ")
	}
	settings.DatabaseSettings.Host = ip[0].String()
	if settings.DatabaseSettings.Database == "" {
		return errors.New("No Database specified ")
	}
	if settings.ServerSettings.Ip == "" {
		return errors.New("No Server Ip specified ")
	}
	if settings.ServerSettings.Port == "" {
		return errors.New("No Server Port specified ")
	}
	if settings.DatabaseSettings.Port == "" {
		settings.DatabaseSettings.Port = "3306"
	}
	if settings.DatabaseSettings.User == "" {
		settings.DatabaseSettings.User = "root"
	}
	if settings.DatabaseSettings.Password == "" {
		settings.DatabaseSettings.Password = "root"
	}
	return nil
}

func testConnection(settings Settings) error {
	dataSourceName := settings.DatabaseSettings.User + ":" +
		settings.DatabaseSettings.Password + "@" +
		"tcp(" + settings.DatabaseSettings.Host + ":" + settings.DatabaseSettings.Port + ")/" +
		settings.DatabaseSettings.Database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
