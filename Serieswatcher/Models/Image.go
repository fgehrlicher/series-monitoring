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
	ImageDummyProviderPath = "Ressources/dummyProvider.png"
	ImageDummyEpisodePath  = "Ressources/dummyEpisode.png"
	ImageDummySeriesPath   = "Ressources/dummySeries.png"
	ImageEpisode           = 1
	ImageProvider          = 2
	ImageSeries            = 3
)

type Image struct {
	OriginURL    string
	RelativePath string
	AbsolutePath string
	ID           int
	Data         []byte
	Settings     Config.Settings
}

func (imageStruct *Image) LoadImageFromFile(imageType int) error {
	if imageStruct.RelativePath == "" {
		errors.New("No Image Path")
	}
	imageStruct.AbsolutePath = path.Join(imageStruct.Settings.ServerSettings.ImagePath, imageStruct.RelativePath)
	if !imageStruct.Exists() {
		dummyImagePath, dummyImagePathError := getDummyImagePath(imageType)
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
	imageData, err := ioutil.ReadFile(imageStruct.AbsolutePath)
	if err != nil {
		return err
	}
	imageStruct.Data = imageData
	return nil
}

func (imageStruct *Image) Create(data io.ReadCloser) error {
	if imageStruct.AbsolutePath == "" {
		imageStruct.AbsolutePath = path.Join(imageStruct.Settings.ServerSettings.ImagePath, imageStruct.RelativePath)
	}
	file, err := os.Create(imageStruct.AbsolutePath)
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
	if imageStruct.AbsolutePath == "" {
		imageStruct.AbsolutePath = path.Join(imageStruct.Settings.ServerSettings.ImagePath, imageStruct.RelativePath)
	}
	_, err := os.Stat(imageStruct.AbsolutePath)
	return err == nil
}

func (imageStruct *Image) Delete() error {
	if imageStruct.AbsolutePath == "" {
		imageStruct.AbsolutePath = path.Join(imageStruct.Settings.ServerSettings.ImagePath, imageStruct.RelativePath)
	}
	err := os.Remove(imageStruct.AbsolutePath)
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

const ImageBasicSelectQuery = "SELECT " + ImageTableName + ".`id`, " + ImageTableName + ".`OriginUrl`, " + ImageTableName + ".`Path` FROM " + ImageTableName

type ImageRepository struct {
	Db *sql.DB
}

func (repository *ImageRepository) Persist(image Image) error {
	query := "INSERT INTO " + ImageTableName + " (OriginUrl, Path) VALUES (?, ?)"
	_, err := repository.Db.Exec(
		query,
		image.OriginURL,
		image.RelativePath,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repository *ImageRepository) GetByPath(path string) (Image, error) {
	var image Image
	query := ImageBasicSelectQuery + " WHERE " + ImageTableName + ".`Path` = ?"
	err := repository.Db.QueryRow(query, path).Scan(&image.ID, &image.OriginURL, &image.RelativePath)
	return image, err
}
