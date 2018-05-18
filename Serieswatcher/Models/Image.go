package Models

import (
	"github.com/murlokswarm/errors"
	"io/ioutil"
	"os"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Config"
	"path"
	"path/filepath"
	"io"
	"database/sql"
)

const (
	ImageTableName         = "Image"
	ImageBasicSelectQuery  = "SELECT " + ImageTableName + ".`id`, " + ImageTableName + ".`OriginUrl`, " + ImageTableName + ".`Path` FROM " + ImageTableName
	ImageDummyProviderPath = "Ressources/dummyProvider.png"
	ImageDummyEpisodePath  = "Ressources/dummyEpisode.png"
	ImageDummySeriesPath   = "Ressources/dummySeries.png"
	ImageEpisode           = 1
	ImageProvider          = 2
	ImageSeries            = 3
)

type ImageRepository struct {
	Db *sql.DB
}
type Image struct {
	ID           int64
	RelativePath string
	OriginURL    string
	Data         []byte
	ImageType    int
}

func (imageStruct *Image) LoadFromFile() error {
	if imageStruct.RelativePath == "" {
		errors.New("No Image Path")
	}
	configuration, _ := Config.GetConfiguration()
	absolutePath := path.Join(configuration.ServerSettings.ImagePath, imageStruct.RelativePath)
	if !imageStruct.Exists() {
		dummyImagePath, dummyImagePathError := getDummyImagePath(imageStruct.ImageType)
		if dummyImagePathError != nil {
			return dummyImagePathError
		}
		imageFile, dummyImageErr := ioutil.ReadFile(dummyImagePath)
		if dummyImageErr != nil {
			return dummyImageErr
		}
		imageStruct.Data = imageFile
		return nil
	}
	imageData, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		return err
	}
	imageStruct.Data = imageData
	return nil
}

func (imageStruct *Image) CreateFile(data io.ReadCloser) error {
	configuration, _ := Config.GetConfiguration()
	absolutePath := path.Join(configuration.ServerSettings.ImagePath, imageStruct.RelativePath)
	if imageStruct.Exists() {
		imageStruct.Delete()
	}
	file, err := os.Create(absolutePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, data)
	if err != nil {
		return err
	}
	return nil
}

func (imageStruct *Image) Exists() bool {
	configuration, _ := Config.GetConfiguration()
	absolutePath := path.Join(configuration.ServerSettings.ImagePath, imageStruct.RelativePath)
	_, err := os.Stat(absolutePath)
	return err == nil
}

func (imageStruct *Image) Delete() error {
	configuration, _ := Config.GetConfiguration()
	absolutePath := path.Join(configuration.ServerSettings.ImagePath, imageStruct.RelativePath)
	err := os.Remove(absolutePath)
	if err != nil {
		return err
	}
	return nil
}

func getDummyImagePath(imageType int) (string, error) {
	var imagePath string
	switch imageType {
	case ImageEpisode:
		imagePath = ImageDummyEpisodePath
	case ImageProvider:
		imagePath = ImageDummyProviderPath
	case ImageSeries:
		imagePath = ImageDummySeriesPath
	}
	return filepath.Abs(imagePath)
}

func (repository *ImageRepository) Persist(image Image) (int64, error) {

	existingImage, _ := repository.GetByPath(image.RelativePath)
	if existingImage.ID > 0 {
		return existingImage.ID, nil
	}

	query := "INSERT INTO " + ImageTableName + " (OriginUrl, Path) VALUES (?, ?)"
	res, err := repository.Db.Exec(
		query,
		image.OriginURL,
		image.RelativePath,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repository *ImageRepository) GetByPath(path string) (*Image, error) {
	var image Image
	query := ImageBasicSelectQuery + " WHERE " + ImageTableName + ".`Path` = ?"
	err := repository.Db.QueryRow(query, path).Scan(&image.ID, &image.OriginURL, &image.RelativePath)
	return &image, err
}

func (repository *ImageRepository) GetById(id int64) (*Image, error) {
	var image Image
	query := ImageBasicSelectQuery + " WHERE " + ImageTableName + ".`id` = ?"
	err := repository.Db.QueryRow(query, id).Scan(&image.ID, &image.OriginURL, &image.RelativePath)
	return &image, err
}
