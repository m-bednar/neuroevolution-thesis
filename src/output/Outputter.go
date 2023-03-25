package output

import (
	"log"
	"os"
	"path"
)

const (
	CHART_FILE_NAME = "chart.png"
	VIDEO_FILE_NAME = "video.avi"
	OUTPUT_DIR_PERM = 0770
)

type Outputter struct {
	chartMaker *ChartMaker
	videoMaker *VideoMaker
}

func CreateOutputPath(outputPath string) {
	if err := os.MkdirAll(outputPath, OUTPUT_DIR_PERM); err != nil {
		log.Fatal(err)
	}
}

func NewOutputter(collector *DataCollector, renderer *Renderer) *Outputter {
	return &Outputter{
		chartMaker: NewChartMaker(collector),
		videoMaker: NewVideoMaker(collector, renderer),
	}
}

func (outputter *Outputter) MakeOutput(outputPath string) {
	CreateOutputPath(outputPath)

	chartPath := path.Join(outputPath, CHART_FILE_NAME)
	videoPath := path.Join(outputPath, VIDEO_FILE_NAME)

	outputter.chartMaker.MakeChart(chartPath)
	outputter.videoMaker.MakeVideoToFile(videoPath)
}
