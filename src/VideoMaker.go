package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"

	"github.com/icza/mjpeg"
)

const (
	FPS          = 14
	JPEG_QUALITY = 98
)

type VideoMaker struct {
	collector *DataCollector
	renderer  *Renderer
}

type Frame *image.RGBA

func NewVideoMaker(collector *DataCollector, renderer *Renderer) *VideoMaker {
	return &VideoMaker{collector, renderer}
}

func (maker *VideoMaker) MakeWritter(filename string) mjpeg.AviWriter {
	var size = int32(maker.renderer.imageSize)
	var writter, err = mjpeg.New(filename, size, size, FPS)
	if err != nil {
		log.Fatal(err)
	}
	return writter
}

func (maker *VideoMaker) MakeVideoToFile(filename string) {
	var writter = maker.MakeWritter(filename)
	var samples = maker.collector.GetCapturedGenerationSamples()
	var total = len(samples) - 1

	for i, sample := range samples {
		maker.AddGenerationSampleFrame(writter, i, sample)
		fmt.Printf("Processing %d/%d\n", i, total)
	}

	if err := writter.Close(); err != nil {
		log.Fatal(err)
	}
}

func (maker *VideoMaker) AddGenerationSampleFrame(writter mjpeg.AviWriter, generation int, sample CapturedGenerationSample) {
	var encoded = maker.EncodeFramesAsync(generation, sample)
	for _, enc := range encoded {
		writter.AddFrame(enc)
	}
}

func (maker *VideoMaker) EncodeFramesAsync(generation int, sample CapturedGenerationSample) [][]byte {
	var encoded = make([][]byte, len(sample.steps))

	LoopAsync(sample.steps, func(index int, _ []Position) {
		var frame = maker.renderer.RenderStep(sample, index)
		encoded[index] = maker.EncodeFrame(frame)
	})

	return encoded
}

func (maker *VideoMaker) EncodeFrame(frame *image.RGBA) []byte {
	var opts = jpeg.Options{Quality: JPEG_QUALITY}
	var buffer = bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(buffer, frame, &opts); err != nil {
		log.Fatal(err)
	}
	return buffer.Bytes()
}
