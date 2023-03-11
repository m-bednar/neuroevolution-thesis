package main

import (
	"log"
	"os"
	"path"
)

const VIDEO_FILE_NAME = "video.avi"
const OUTPUT_DIRECTORY_PERMISSIONS = 0770

type Outputter struct {
	collector  *DataCollector
	chartMaker *ChartMaker
	videoMaker *VideoMaker
}

func CreateOutputPath(outputPath string) {
	var err = os.MkdirAll(outputPath, OUTPUT_DIRECTORY_PERMISSIONS)
	if err != nil {
		log.Fatal(err)
	}
}

func NewOutputter(collector *DataCollector, renderer *Renderer) *Outputter {
	return &Outputter{
		collector:  collector,
		chartMaker: NewChartMaker(),
		videoMaker: NewVideoMaker(renderer),
	}
}

func (outputter *Outputter) MakeOutput(outputPath string) {
	CreateOutputPath(outputPath)
	var videoPath = path.Join(outputPath, VIDEO_FILE_NAME)

	// TODO: Make gourutines

	outputter.chartMaker.MakeChart(outputter.collector)
	outputter.videoMaker.MakeVideoToFile(videoPath, outputter.collector)
}
