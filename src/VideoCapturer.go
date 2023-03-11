package main

import (
	"bytes"
	"image"
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

func (maker *VideoMaker) MakeVideoToFile(filename string, collector *DataCollector) {
	var writter = maker.MakeWritter(filename)
	var frames = maker.MakeFrames(collector.samples)

	// TODO: Use goroutines
	for i := range frames {
		var encoded = maker.EncodeFrame(frames[i])
		writter.AddFrame(encoded)
	}

	if err := writter.Close(); err != nil {
		log.Fatal(err)
	}
}

func (maker *VideoMaker) EncodeFrame(frame *image.RGBA) []byte {
	var opts = jpeg.Options{Quality: JPEG_QUALITY}
	var buffer = bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(buffer, frame, &opts); err != nil {
		log.Fatal(err)
	}
	return buffer.Bytes()
}

func (maker *VideoMaker) MakeFrames(samples []GenerationSample) []Frame {
	var frames = make([]Frame, 0)
	for i, sample := range samples {
		frames = append(frames, maker.CaptureGeneration(i, sample)...)
	}
	return frames
}

func (maker *VideoMaker) CaptureGeneration(generation int, sample GenerationSample) []Frame {
	var steps = make([]Frame, len(sample.steps))
	for i, step := range sample.steps {
		var frame = maker.renderer.RenderStep(generation, step)
		steps[i] = frame
	}
	return steps
}
