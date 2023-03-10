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

func NewOutputter(outputPath string, collector *DataCollector, renderer *Renderer) *Outputter {
	CreateOutputPath(outputPath)
	var videoPath = path.Join(outputPath, VIDEO_FILE_NAME)
	return &Outputter{
		collector:  collector,
		chartMaker: NewChartMaker(),
		videoMaker: NewVideoMaker(videoPath, renderer),
	}
}

func (outputter *Outputter) SaveAll() {
	outputter.videoMaker.SaveVideo()
}
