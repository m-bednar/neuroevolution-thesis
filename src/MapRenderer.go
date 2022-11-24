package main

import (
	"image"
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
)

func drawCircle(context *draw2dimg.GraphicContext, x, y float64) {
	const radius = 3
	context.SetFillColor(color.RGBA{20, 20, 200, 255})
	context.SetLineWidth(0)
	context.BeginPath()
	context.ArcTo(x, y, radius, radius, 0, 2 * math.Pi)
	context.FillStroke()
}

func Render() {
	var dest = image.NewRGBA(image.Rect(0, 0, 700, 700))
	var context = draw2dimg.NewGraphicContext(dest)
	drawCircle(context, 20, 40)	// diameter: 6px, 100x100 = 600x600px + grid 100px = 700x700px
	draw2dimg.SaveToPngFile("hello.png", dest)
}
