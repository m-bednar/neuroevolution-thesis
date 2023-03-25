package output

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"

	"github.com/icza/mjpeg"
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	. "github.com/m-bednar/neuroevolution-thesis/src/utils"
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
	size := int32(maker.renderer.imageSize)
	writter, err := mjpeg.New(filename, size, size, FPS)
	if err != nil {
		log.Fatal(err)
	}
	return writter
}

func (maker *VideoMaker) MakeVideoToFile(filename string) {
	writter := maker.MakeWritter(filename)
	samples := maker.collector.GetCapturedGenerationSamples()
	total := len(samples) - 1

	for i, sample := range samples {
		maker.AddGenerationSampleFrame(writter, i, sample)
		fmt.Printf("Processing %d/%d\n", i, total)
	}

	if err := writter.Close(); err != nil {
		log.Fatal(err)
	}
}

func (maker *VideoMaker) AddGenerationSampleFrame(writter mjpeg.AviWriter, generation int, sample CapturedGenerationSample) {
	encoded := maker.EncodeFramesAsync(generation, sample)
	for _, enc := range encoded {
		writter.AddFrame(enc)
	}
}

func (maker *VideoMaker) EncodeFramesAsync(generation int, sample CapturedGenerationSample) [][]byte {
	encoded := make([][]byte, len(sample.steps))

	ConcurrentLoop(sample.steps, func(index int, _ []Position) {
		frame := maker.renderer.RenderStep(sample, index)
		encoded[index] = maker.EncodeFrame(frame)
	})

	return encoded
}

func (maker *VideoMaker) EncodeFrame(frame *image.RGBA) []byte {
	opts := jpeg.Options{Quality: JPEG_QUALITY}
	buffer := bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(buffer, frame, &opts); err != nil {
		log.Fatal(err)
	}
	return buffer.Bytes()
}
