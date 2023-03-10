package main

import (
	"bytes"
	"image/jpeg"
	"log"

	"github.com/icza/mjpeg"
)

const (
	FPS          = 15
	JPEG_QUALITY = 95
)

type VideoMaker struct {
	renderer *Renderer
	writter  mjpeg.AviWriter
}

func NewVideoMaker(filename string, renderer *Renderer) *VideoMaker {
	var size = int32(renderer.imageSize)
	var writter, err = mjpeg.New(filename, size, size, FPS)

	if err != nil {
		log.Fatal(err)
	}

	return &VideoMaker{
		renderer: renderer,
		writter:  writter,
	}
}

func (maker *VideoMaker) MakeVideo(collector *DataCollector) {
	// TODO
}

func (maker *VideoMaker) CaptureScene(generation int, population []*Microbe) {
	var opts = jpeg.Options{Quality: JPEG_QUALITY}
	var img = maker.renderer.RenderScene(generation, population)

	var buffer = bytes.NewBuffer([]byte{})
	var err = jpeg.Encode(buffer, img, &opts)

	if err != nil {
		log.Fatal(err)
	}
	maker.writter.AddFrame(buffer.Bytes())
}

func (maker *VideoMaker) SaveVideo() {
	var err = maker.writter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
