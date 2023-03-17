package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"sync"

	"github.com/icza/mjpeg"
)

const (
	FPS            = 12
	JPEG_QUALITY   = 98
	CAPTURE_MARGIN = 10
)

type VideoMaker struct {
	renderer *Renderer
}

type Frame *image.RGBA

func NewVideoMaker(renderer *Renderer) *VideoMaker {
	return &VideoMaker{renderer}
}

func (maker *VideoMaker) MakeWritter(filename string) mjpeg.AviWriter {
	var size = int32(maker.renderer.imageSize)
	var writter, err = mjpeg.New(filename, size, size, FPS)
	if err != nil {
		log.Fatal(err)
	}
	return writter
}

func (maker *VideoMaker) MakeVideoToFile(filename string, collector *DataCollector) {
	var writter = maker.MakeWritter(filename)
	var samples = collector.GetCapturedGenerationSamples()
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
	var wg = sync.WaitGroup{}
	var total = len(sample.steps)
	var encoded = make([][]byte, total)

	wg.Add(total)
	for i := range sample.steps {
		go func(j int) {
			var frame = maker.renderer.RenderStep(sample, j)
			encoded[j] = maker.EncodeFrame(frame)
			wg.Done()
		}(i)
	}
	wg.Wait()

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
