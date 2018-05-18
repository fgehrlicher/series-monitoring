package Models

import (
	"time"
	"database/sql"
	"github.com/murlokswarm/errors"
	"strconv"
)

const (
	EpisodeTableName        = "Episode"
	EpisodeBasicSelectQuery = "SELECT " +
		EpisodeTableName + ".`id`, " +
		EpisodeTableName + ".`Series_id`, " +
		EpisodeTableName + ".`Image_id`, " +
		EpisodeTableName + ".`Episode`, " +
		EpisodeTableName + ".`Season`, " +
		EpisodeTableName + ".`Title`, " +
		EpisodeTableName + ".`Description`, " +
		EpisodeTableName + ".`ReleaseDate` FROM " +
		EpisodeTableName
)

type Episode struct {
	ID          int64
	Series      *Series `json:"-"`
	Image       *Image  `json:"-"`
	Episode     int
	Season      int
	Title       string
	Description string
	ReleaseDate time.Time
}

type EpisodeRepository struct {
	Db *sql.DB
}

func (repository *EpisodeRepository) Persist(episode Episode) (int64, error) {
	imageId := sql.NullInt64{Int64: episode.Image.ID}

	if episode.Image.ID > 0 {
		imageId.Valid = true
	}

	query := "INSERT INTO `" + EpisodeTableName +
		"` (`Series_id`, `Image_id`, `Episode`, `Season`, `Title`, `Description`, `ReleaseDate`)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := repository.Db.Exec(
		query,
		episode.Series.ID,
		imageId,
		episode.Episode,
		episode.Season,
		episode.Title,
		episode.Description,
		episode.ReleaseDate.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (repository *EpisodeRepository) Update(episode Episode) error {
	return nil
}

func (repository *EpisodeRepository) GetAllBySeries(series Series, resolveRelations bool) ([]Episode, error) {
	var episodes []Episode

	seriesId := series.ID
	if seriesId == 0 {
		return nil, errors.New("Series is not persisted")
	}

	row, err := repository.Db.Query(EpisodeBasicSelectQuery+" WHERE "+EpisodeTableName+".`Series_id` = ?", seriesId)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	for row.Next() {
		SeriesID := sql.NullInt64{}
		ImageId := sql.NullInt64{}
		episode := Episode{}

		row.Scan(
			&episode.ID,
			&SeriesID,
			&ImageId,
			&episode.Episode,
			&episode.Season,
			&episode.Title,
			&episode.Description,
			&episode.ReleaseDate,
		)

		if resolveRelations {
			if SeriesID.Valid {
				seriesRepository := SeriesRepository{repository.Db}
				series, err := seriesRepository.GetById(SeriesID.Int64, false)
				if err != nil {
					return nil, err
				}
				episode.Series = series

			}
			if ImageId.Valid {
				imageRepository := ImageRepository{repository.Db}
				image, err := imageRepository.GetById(ImageId.Int64)
				if err != nil {
					return nil, err
				}
				episode.Image = image
			}
		}

		episodes = append(episodes, episode)
	}

	return episodes, nil
}

func (repository *EpisodeRepository) GetAllBySeriesAndSeason(series Series, season int, resolveRelations bool) ([]Episode, error) {
	var episodes []Episode

	seriesId := series.ID
	if seriesId == 0 {
		return nil, errors.New("Series is not persisted")
	}

	row, err := repository.Db.Query(
		EpisodeBasicSelectQuery + " WHERE " + EpisodeTableName + ".`Series_id` = " + strconv.Itoa(int(seriesId)) + " AND " + EpisodeTableName + ".`Season` = " + strconv.Itoa(int(season)))
	if err != nil {
		return nil, err
	}

	defer row.Close()
	for row.Next() {
		SeriesID := sql.NullInt64{}
		ImageId := sql.NullInt64{}
		episode := Episode{}

		row.Scan(
			&episode.ID,
			&SeriesID,
			&ImageId,
			&episode.Episode,
			&episode.Season,
			&episode.Title,
			&episode.Description,
			&episode.ReleaseDate,
		)

		if resolveRelations {
			if SeriesID.Valid {
				seriesRepository := SeriesRepository{repository.Db}
				series, err := seriesRepository.GetById(SeriesID.Int64, false)
				if err != nil {
					return nil, err
				}
				episode.Series = series

			}
			if ImageId.Valid {
				imageRepository := ImageRepository{repository.Db}
				image, err := imageRepository.GetById(ImageId.Int64)
				if err != nil {
					return nil, err
				}
				episode.Image = image
			}
		}

		episodes = append(episodes, episode)
	}

	return episodes, nil
}

func (repository *EpisodeRepository) GetOneBySeriesAndSeasonAndEpisode(series Series, seasonId int, episodeId int, resolveRelations bool) (*Episode, error) {
	var episode Episode
	query := EpisodeBasicSelectQuery + " WHERE `Series_id` = " + strconv.Itoa(int(series.ID)) + " AND `Season` = " + strconv.Itoa(int(seasonId)) + " AND `Episode` = " + strconv.Itoa(int(episodeId)) + " LIMIT 1;"
	SeriesID := sql.NullInt64{}
	ImageId := sql.NullInt64{}

	err := repository.Db.QueryRow(query).Scan(
		&episode.ID,
		&SeriesID,
		&ImageId,
		&episode.Episode,
		&episode.Season,
		&episode.Title,
		&episode.Description,
		&episode.ReleaseDate)
	if err != nil {
		return nil, err
	}
	if resolveRelations {
		if SeriesID.Valid {
			seriesRepository := SeriesRepository{repository.Db}
			series, err := seriesRepository.GetById(SeriesID.Int64, false)
			if err != nil {
				return nil, err
			}
			episode.Series = series

		}
		if ImageId.Valid {
			imageRepository := ImageRepository{repository.Db}
			image, err := imageRepository.GetById(ImageId.Int64)
			if err != nil {
				return nil, err
			}
			episode.Image = image
		}
	}
	return &episode, nil
}

func (repository *EpisodeRepository) GetById(id int64, resolveRelations bool) (*Episode, error) {
	var episode Episode
	query := EpisodeBasicSelectQuery + " WHERE " + EpisodeTableName + ".`id` = ?"
	SeriesID := sql.NullInt64{}
	ImageId := sql.NullInt64{}

	err := repository.Db.QueryRow(query, id).Scan(
		&episode.ID,
		&SeriesID,
		&ImageId,
		&episode.Episode,
		&episode.Season,
		&episode.Title,
		&episode.Description,
		&episode.ReleaseDate)
	if err != nil {
		return nil, err
	}
	if resolveRelations {
		if SeriesID.Valid {
			seriesRepository := SeriesRepository{repository.Db}
			series, err := seriesRepository.GetById(SeriesID.Int64, false)
			if err != nil {
				return nil, err
			}
			episode.Series = series

		}
		if ImageId.Valid {
			imageRepository := ImageRepository{repository.Db}
			image, err := imageRepository.GetById(ImageId.Int64)
			if err != nil {
				return nil, err
			}
			episode.Image = image
		}
	}
	return &episode, err
}

func (repository *EpisodeRepository) GetLatestBySeries(series Series) (*Episode, error) {
	var episode Episode
	query := EpisodeBasicSelectQuery + " WHERE Episode.`Series_id` = ? ORDER BY Episode.`Season` DESC, Episode.`Episode` DESC LIMIT 1; "
	SeriesID := sql.NullInt64{}
	ImageId := sql.NullInt64{}

	err := repository.Db.QueryRow(query, series.ID).Scan(
		&episode.ID,
		&SeriesID,
		&ImageId,
		&episode.Episode,
		&episode.Season,
		&episode.Title,
		&episode.Description,
		&episode.ReleaseDate)
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

func (repository *EpisodeRepository) GetAllNewEpisodes(series Series) ([]Episode, error) {
	var episodes []Episode

	seriesId := series.ID
	if seriesId == 0 {
		return nil, errors.New("Series is not persisted")
	}

	if series.WatchPointer == nil {
		return nil, errors.New("No WatchPointer set")
	}

	row, err := repository.Db.Query(
		EpisodeBasicSelectQuery + " WHERE " + EpisodeTableName + ".`Series_id` = " + strconv.Itoa(int(seriesId)) +
			" AND " + EpisodeTableName + ".`Season` >= " + strconv.Itoa(int(series.WatchPointer.Season)) +
			" AND " + EpisodeTableName + ".`Episode` > " + strconv.Itoa(int(series.WatchPointer.Episode)) +
			" ORDER BY Episode.`Season` ASC, Episode.`Episode` ASC;",
	)
	if err != nil {
		return nil, err
	}

	defer row.Close()
	for row.Next() {
		SeriesID := sql.NullInt64{}
		ImageId := sql.NullInt64{}
		episode := Episode{}

		row.Scan(
			&episode.ID,
			&SeriesID,
			&ImageId,
			&episode.Episode,
			&episode.Season,
			&episode.Title,
			&episode.Description,
			&episode.ReleaseDate,
		)

		episodes = append(episodes, episode)
	}
	return episodes, nil
}
