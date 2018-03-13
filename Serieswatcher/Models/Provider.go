package Models

import (
	"database/sql"
)

const (
	ProviderTableName = "Provider"
)

type Provider struct {
	ImageID   int
	ImagePath string
	Name      string
	ID        int
}

type ProviderRepository struct {
	Db *sql.DB
}

func (repository *ProviderRepository) GetAll() ([]Provider, error) {
	var providers []Provider
	query := "SELECT " +
		ProviderTableName + ".`id`, " +
		ProviderTableName + ".`Name`, " +
		ProviderTableName + ".`Image_id`," +
		ImageTableName + " .`Path` " +
		"FROM " + ProviderTableName + " LEFT JOIN " + ImageTableName +
		" ON " + ProviderTableName + ".`Image_id` = " + ImageTableName + ".`id`"
	row, err := repository.Db.Query(query)
	if err != nil {
		return providers, err
	}
	defer row.Close()
	for row.Next() {
		r := Provider{}
		row.Scan(&r.ID, &r.Name, &r.ImageID, &r.ImagePath)
		providers = append(providers, r)
	}
	return providers, err
}

func (repository *ProviderRepository) GetByName(name string) (Provider, error) {
	var provider Provider
	query := "SELECT " +
		ProviderTableName + ".`id`, " +
		ProviderTableName + ".`Name`, " +
		ProviderTableName + ".`Image_id`," +
		ImageTableName + " .`Path` " +
		"FROM " + ProviderTableName + " LEFT JOIN " + ImageTableName +
		" ON " + ProviderTableName + ".`Image_id` = " + ImageTableName + ".`id` " +
		"WHERE " + ProviderTableName + ".`Name` = ? LIMIT 1"
	err := repository.Db.QueryRow(query, name).Scan(&provider.ID, &provider.Name, &provider.ImageID, &provider.ImagePath)
	return provider, err
}
