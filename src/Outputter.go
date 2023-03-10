package main

import (
	"log"
	"os"
	"path"
)

const VIDEO_FILE_NAME = "video.avi"
const OUTPUT_DIRECTORY_PERMISSIONS = 0770

type Outputter struct {
	videoCapturer *VideoCapturer
}

func CreateOutputPath(outputPath string) {
	var err = os.MkdirAll(outputPath, OUTPUT_DIRECTORY_PERMISSIONS)
	if err != nil {
		log.Fatal(err)
	}
}

func NewOutputter(outputPath string, renderer *Renderer) *Outputter {
	CreateOutputPath(outputPath)
	var videoPath = path.Join(outputPath, VIDEO_FILE_NAME)
	return &Outputter{
		videoCapturer: NewVideoCapturer(videoPath, renderer),
	}
}

func (outputter *Outputter) CaptureStep(generation int, population []*Microbe) {
	outputter.videoCapturer.CaptureScene(generation, population)
}

func (outputter *Outputter) SaveAll() {
	outputter.videoCapturer.SaveVideo()
}
