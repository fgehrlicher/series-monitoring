package Models

import (
	"database/sql"
	"errors"
	"strconv"
	"encoding/json"
	"net/url"
)

const (
	SeriesTableName        = "Series"
	SeriesBasicSelectQuery = "SELECT " +
		SeriesTableName + ".`id`, " +
		SeriesTableName + ".`Title`, " +
		SeriesTableName + ".`ProviderURL`, " +
		SeriesTableName + ".`Image_id`, " +
		SeriesTableName + ".`Episode_id` " +
		"FROM " + SeriesTableName
)

type Series struct {
	ID                int64
	Title             string
	ProviderURL       string
	UnwatchedEpisodes []Episode
	WatchPointer      *Episode
	Image             *Image
}

func (series Series) MarshalJSON() (b []byte, e error) {
	return json.Marshal(struct {
		ID                int64     `json:"id"`
		Title             string    `json:"title"`
		ImageUrl          string    `json:"image_path"`
		ProviderURL       string    `json:"provider_url"`
		UnwatchedEpisodes []Episode `json:"unwatched_episodes"`
		WatchPointer      *Episode  `json:"current_episode"`
	}{
		ID:                series.ID,
		Title:             series.Title,
		ImageUrl:          "series/" + url.PathEscape(series.Title) + "/image/",
		ProviderURL:       series.ProviderURL,
		UnwatchedEpisodes: series.UnwatchedEpisodes,
		WatchPointer:      series.WatchPointer,
	})
}

type SeriesRepository struct {
	Db *sql.DB
}

func (repository *SeriesRepository) GetByName(title string, resolveRelations bool) (*Series, error) {
	return repository.getByIdentifier("Title", title, resolveRelations)
}

func (repository *SeriesRepository) GetByProviderURL(providerURL string, resolveRelations bool) (*Series, error) {
	return repository.getByIdentifier("ProviderURL", providerURL, resolveRelations)
}

func (repository *SeriesRepository) GetAll(resolveRelations bool) ([]Series, error) {
	var series []Series
	var imageRepository ImageRepository
	var episodeRepository EpisodeRepository
	var imageId int64
	var episodeId int64
	if resolveRelations {
		imageRepository.Db = repository.Db
		episodeRepository.Db = repository.Db
	}
	query := SeriesBasicSelectQuery
	row, err := repository.Db.Query(query)
	if err != nil {
		return series, err
	}
	defer row.Close()
	for row.Next() {
		r := Series{}
		row.Scan(&r.ID, &r.Title, &r.ProviderURL, &imageId, &episodeId)
		if resolveRelations {
			if imageId > 0 {
				imageStruct, err := imageRepository.GetById(imageId)
				if err != nil {
					return nil, err
				}
				imageStruct.ImageType = ImageProvider
				r.Image = imageStruct
			}
			if episodeId > 0 {
				episode, err := episodeRepository.GetById(episodeId, resolveRelations)
				if err != nil {
					return nil, err
				}
				r.WatchPointer = episode
			}
		}
		series = append(series, r)
	}
	return series, err
}

func (repository *SeriesRepository) getByIdentifier(identifier string, identifierValue string, resolveRelations bool) (*Series, error) {
	series := Series{}
	query := SeriesBasicSelectQuery + " WHERE " + SeriesTableName + ".`" + identifier + "` = ?"
	var (
		watchpointerId sql.NullInt64
		imageId        sql.NullInt64
	)
	err := repository.Db.QueryRow(query, identifierValue).Scan(&series.ID, &series.Title, &series.ProviderURL, &imageId, &watchpointerId)
	if err != nil {
		return nil, err
	}
	if resolveRelations {
		if imageId.Valid {
			imageRepository := ImageRepository{repository.Db}
			imageStruct, err := imageRepository.GetById(imageId.Int64)
			if err != nil {
				return nil, err
			}
			imageStruct.ImageType = ImageProvider
			series.Image = imageStruct
		}
		if watchpointerId.Valid {
			episodeRepository := EpisodeRepository{repository.Db}
			episode, err := episodeRepository.GetById(watchpointerId.Int64, resolveRelations)
			if err != nil {
				return nil, err
			}
			series.WatchPointer = episode
		}
	}
	return &series, nil
}

func (repository *SeriesRepository) GetById(id int64, resolveRelations bool) (*Series, error) {
	series := Series{}
	query := SeriesBasicSelectQuery + " WHERE " + SeriesTableName + ".`id` = ?"
	var (
		watchpointerId int64
		imageId        int64
	)
	err := repository.Db.QueryRow(query, id).Scan(&series.ID, &series.Title, &series.ProviderURL, &imageId, &watchpointerId)
	if err != nil {
		return nil, err
	}
	if resolveRelations {
		if imageId > 0 {
			imageRepository := ImageRepository{repository.Db}
			imageStruct, err := imageRepository.GetById(imageId)
			if err != nil {
				return nil, err
			}
			imageStruct.ImageType = ImageProvider
			series.Image = imageStruct
		}
		if watchpointerId > 0 {
			episodeRepository := EpisodeRepository{repository.Db}
			episode, err := episodeRepository.GetById(watchpointerId, resolveRelations)
			if err != nil {
				return nil, err
			}
			series.WatchPointer = episode
		}
	}
	return &series, nil
}

func (repository *SeriesRepository) Persist(series Series) (int64, error) {
	query := "INSERT INTO " + SeriesTableName + " (ProviderURL, Title, Image_id) VALUES (?, ?, ?)"
	result, err := repository.Db.Exec(query, series.ProviderURL, series.Title, series.Image.ID)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (repository *SeriesRepository) UpdateWatchPointer(series *Series) error {
	if !(series.ID > 0 && series.WatchPointer.ID > 0) {
		return errors.New("WatchPointer or Series are not persisted")
	}
	query := "UPDATE " + SeriesTableName + " SET Episode_id = " +
		strconv.Itoa(int(series.WatchPointer.ID)) +
		" WHERE id = " + strconv.Itoa(int(series.ID))
	_, err := repository.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
