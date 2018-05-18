package Models

import (
	"database/sql"
)

const (
	ProviderTableName = "Provider"
)

type Provider struct {
	ID   int
	Name string
	Image *Image `json:"-"`
}

type ProviderRepository struct {
	Db *sql.DB
}

func (repository *ProviderRepository) GetAll(resolveRelations bool) ([]Provider, error) {
	var providers []Provider
	var imageRepository ImageRepository
	if resolveRelations {
		imageRepository.Db = repository.Db
	}
	query := "SELECT " +
		ProviderTableName + ".`id`, " +
		ProviderTableName + ".`Name`, " +
		ProviderTableName + ".`Image_id` " +
		"FROM " + ProviderTableName
	row, err := repository.Db.Query(query)
	if err != nil {
		return providers, err
	}
	defer row.Close()
	var imageId int64
	for row.Next() {
		r := Provider{}
		row.Scan(&r.ID, &r.Name, &imageId)
		if resolveRelations {
			image, err := imageRepository.GetById(imageId)
			if err != nil {
				return providers, err
			}
			r.Image = image
		}
		providers = append(providers, r)
	}
	return providers, err
}

func (repository *ProviderRepository) GetByName(name string, resolveRelations bool) (Provider, error) {
	var provider Provider
	var imageRepository ImageRepository
	var imageId int64
	if resolveRelations {
		imageRepository.Db = repository.Db
	}
	query := "SELECT " +
		ProviderTableName + ".`id`, " +
		ProviderTableName + ".`Name`, " +
		ProviderTableName + ".`Image_id` " +
		"FROM " + ProviderTableName + " WHERE " + ProviderTableName + ".`Name` = ? LIMIT 1"
	err := repository.Db.QueryRow(query, name).Scan(&provider.ID, &provider.Name, &imageId)
	if err != nil {
		return provider, err
	}
	if resolveRelations {
		image, err := imageRepository.GetById(imageId)
		if err != nil {
			return provider, err
		}
		image.ImageType = ImageProvider
		provider.Image = image
	}
	return provider, err
}
