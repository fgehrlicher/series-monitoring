package Models

import "database/sql"

const (
	SeriesTableName        = "Series"
	SeriesBasicSelectQuery = "SELECT " +
		SeriesTableName + ".`id`, " +
		SeriesTableName + ".`ProviderURL`, " +
		SeriesTableName + ".`Title`, " +
		SeriesTableName + ".`Image_id`, " +
		ImageTableName + " .`Path` " +
		"FROM " + SeriesTableName + " LEFT JOIN " + ImageTableName +
		" ON " + SeriesTableName + ".`Image_id` = " + ImageTableName + ".`id` "
)

type Series struct {
	ID             int
	ImageID        int
	ImagePath      string
	ImageOriginUrl string `json:"-"`
	ProviderURL    string
	Title          string
}

func (repository *SeriesRepository) Scan(row *sql.Rows) Series {
	series := Series{}
	row.Scan(&series.ID, &series.ProviderURL, &series.Title, &series.ImageID, &series.ImagePath)
	return series
}

type SeriesRepository struct {
	Db *sql.DB
}

func (repository *SeriesRepository) GetAll() ([]Series, error) {
	var series []Series
	query := SeriesBasicSelectQuery
	row, err := repository.Db.Query(query)
	if err != nil {
		return series, err
	}
	defer row.Close()
	for row.Next() {
		series = append(series, repository.Scan(row))
	}
	return series, err
}

func (repository *SeriesRepository) GetByName(title string) (Series, error) {
	var series Series
	query := SeriesBasicSelectQuery + " WHERE " + SeriesTableName + ".`Title` = ?"
	err := repository.Db.QueryRow(query, title).Scan(&series.ID, &series.ProviderURL, &series.Title, &series.ImageID, &series.ImagePath)
	return series, err
}

func (repository *SeriesRepository) GetByProviderURL(providerURL string) (Series, error) {
	var series Series
	query := SeriesBasicSelectQuery + " WHERE " + SeriesTableName + ".`ProviderURL` = ?"
	err := repository.Db.QueryRow(query, providerURL).Scan(&series.ID, &series.ProviderURL, &series.Title, &series.ImageID, &series.ImagePath)
	return series, err
}

func (repository *SeriesRepository) Persist(series Series) error {
	//@TODO refactor
	imageRepository := ImageRepository{repository.Db}
	image, err := imageRepository.GetByPath(series.ImagePath)
	if err != nil {
		if err == sql.ErrNoRows {
			err = imageRepository.Persist(
				Image{
					OriginURL:    series.ImageOriginUrl,
					RelativePath: series.ImagePath,
				},
			)
			if err != nil {
				return err
			}
			image, err = imageRepository.GetByPath(series.ImagePath)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	query := "INSERT INTO " + SeriesTableName + " (ProviderURL, Title, Image_id) VALUES (?, ?, ?)"
	_, err = repository.Db.Exec(
		query,
		series.ProviderURL,
		series.Title,
		image.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
