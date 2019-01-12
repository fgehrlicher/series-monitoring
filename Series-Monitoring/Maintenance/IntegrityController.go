package Maintenance

import (
		"github.com/fgehrlicher/series-monitoring/Series-Monitoring/Models"
	"net/http"
	"github.com/fgehrlicher/series-monitoring/Series-Monitoring/Parser"
	"os"
	"github.com/fgehrlicher/series-monitoring/Series-Monitoring/Config"
	"path"
	"strings"
)

type IntegrityController struct {
	report []string
}

func NewIntegrityController() *IntegrityController {
	return &IntegrityController{}
}

func (this *IntegrityController) CheckMetaDataIntegrity(episodes []Models.Episode) (int, string) {
	this.resetReports()
	UpdateCount := 0



	this.writeReport("Check Meta Data Integrity")
	return UpdateCount, strings.Join(this.report, " \n")
}

func (this *IntegrityController) CheckImageIntegrity(images []Models.Image) (int, string) {
	this.resetReports()
	downloadCount := 0

	configuration, _ := Config.GetConfiguration()
	absoluteCoverImagePath := path.Join(configuration.ServerSettings.ImagePath, Parser.CoverImagesPath)
	absoluteEpisodeImagePath := path.Join(configuration.ServerSettings.ImagePath, Parser.EpisodeImagesPath)

	if _, err := os.Stat(absoluteCoverImagePath); os.IsNotExist(err) {
		os.Mkdir(absoluteCoverImagePath, 0777)
	}
	if _, err := os.Stat(absoluteEpisodeImagePath); os.IsNotExist(err) {
		os.Mkdir(absoluteEpisodeImagePath, 0777)
	}

	for _, image := range images {
		if image.Exists() {
			continue
		}
		response, err := http.Get(image.OriginURL)
		if err != nil {
			this.log("Cant retrieve Image: " + image.OriginURL)
			continue
		}
		err = image.CreateFile(response.Body)
		response.Body.Close()
		if err != nil {
			this.log("Cant create Image: " + image.OriginURL)
			continue
		}
		this.log("Downloaded Image: " + image.OriginURL)
		downloadCount ++
	}

	this.writeReport("Check Image Integrity")
	return downloadCount, strings.Join(this.report, " \n")
}

func (this *IntegrityController) resetReports() {
	this.report = nil
}

func (this *IntegrityController) log(message string) {
	this.report = append(this.report, message)
}

func (this *IntegrityController) writeReport(task string) {
	wrapInSuperFancyBox(this.report, task)
}
