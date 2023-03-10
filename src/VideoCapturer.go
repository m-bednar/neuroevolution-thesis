package main

import (
	"bytes"
	"image/jpeg"
	"log"

	"github.com/icza/mjpeg"
)

const (
	FPS          = 15
	JPEG_QUALITY = 90
)

type VideoCapturer struct {
	renderer *Renderer
	writter  mjpeg.AviWriter
}

func NewVideoCapturer(filename string, renderer *Renderer) *VideoCapturer {
	var size = int32(renderer.imageSize)
	var writter, err = mjpeg.New(filename, size, size, FPS)

	if err != nil {
		log.Fatal(err)
	}

	return &VideoCapturer{
		renderer: renderer,
		writter:  writter,
	}
}

func (capturer *VideoCapturer) CaptureScene(generation int, population []*Microbe) {
	var opts = jpeg.Options{Quality: JPEG_QUALITY}
	var img = capturer.renderer.RenderScene(generation, population)

	var buffer = bytes.NewBuffer([]byte{})
	var err = jpeg.Encode(buffer, img, &opts)

	if err != nil {
		log.Fatal(err)
	}
	capturer.writter.AddFrame(buffer.Bytes())
}

func (capturer *VideoCapturer) SaveVideo() {
	var err = capturer.writter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
