package Models

import (
	"github.com/murlokswarm/errors"
	"io/ioutil"
	"os"
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher/Config"
	"path"
	"path/filepath"
)

const (
	ImageTableName         = "Image"
	ImageDummyProviderPath = "Ressources/dummyProvider.png"
	ImageDummyEpisodePath  = "Ressources/dummyEpisode.png"
	ImageDummySeriesPath   = "Ressources/dummySeries.png"
)

type Image struct {
	OriginURL    string
	RelativePath string
	ID           int
	Data         []byte
	Settings     Config.Settings
}

func (imageStruct *Image) LoadImageFromFile() error {
	if imageStruct.RelativePath == "" {
		errors.New("No Image Path")
	}

	imageFile, err := ioutil.ReadFile(path.Join(imageStruct.Settings.ServerSettings.ImagePath, imageStruct.RelativePath))
	if err != nil {
		if os.IsNotExist(err) {
			dummyImagePath, dummyImagePathError := filepath.Abs(ImageDummyProviderPath)
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
		return err
	}
	imageStruct.Data = imageFile
	return nil
}
