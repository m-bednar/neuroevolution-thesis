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
	FPS          = 12
	JPEG_QUALITY = 90
)

type VideoMaker struct {
	renderer *Renderer
}

type Frame *image.RGBA

func NewVideoMaker(renderer *Renderer) *VideoMaker {
	return &VideoMaker{renderer: renderer}
}

func (maker *VideoMaker) MakeWritter(filename string) mjpeg.AviWriter {
	var size = int32(maker.renderer.imageSize)
	var writter, err = mjpeg.New(filename, size, size, FPS)
	if err != nil {
		log.Fatal(err)
	}
	return writter
}

func (maker *VideoMaker) MakeVideoToFile(filename string, captureModifier int, collector *DataCollector) {
	var writter = maker.MakeWritter(filename)
	var count = (len(collector.samples) - 1) / captureModifier

	for i, sample := range collector.samples {
		if i%captureModifier == 0 {
			maker.AddGenerationSampleFrame(writter, i, sample)
			fmt.Printf("Processing %d/%d\n", i/captureModifier, count)
		}
	}

	if err := writter.Close(); err != nil {
		log.Fatal(err)
	}
}

func (maker *VideoMaker) AddGenerationSampleFrame(writter mjpeg.AviWriter, generation int, sample GenerationSample) {
	var encoded = maker.EncodeFramesAsync(generation, sample)
	for _, enc := range encoded {
		writter.AddFrame(enc)
	}
}

func (maker *VideoMaker) EncodeFramesAsync(generation int, sample GenerationSample) [][]byte {
	var wg = sync.WaitGroup{}
	var count = len(sample.steps)
	var encoded = make([][]byte, count)

	wg.Add(count)
	for i, s := range sample.steps {
		go func(ind int, step StepSample) {
			var frame = maker.renderer.RenderStep(generation, step)
			encoded[ind] = maker.EncodeFrame(frame)
			wg.Done()
		}(i, s)
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
