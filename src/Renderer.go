package main

import (
	"image"
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
)

const TILE_DISPLAY_SIZE = 8

type Renderer struct {
	enviroment *Enviroment
	context    *draw2dimg.GraphicContext
}

func NewRenderer(enviroment *Enviroment) *Renderer {
	var size = enviroment.size * TILE_DISPLAY_SIZE
	var image = image.NewRGBA(image.Rect(0, 0, size, size))
	var context = draw2dimg.NewGraphicContext(image)

	context.SetFillColor(color.RGBA{20, 20, 200, 255})
	context.SetLineWidth(0)

	return &Renderer{
		enviroment: enviroment,
		context:    context,
	}
}

func (renderer *Renderer) DrawCircle(x, y float64) {
	const radius = TILE_DISPLAY_SIZE / 2
	const circle = 2 * math.Pi

	renderer.context.BeginPath()
	renderer.context.ArcTo(x, y, radius, radius, 0, circle)
	renderer.context.FillStroke()
}

func (renderer *Renderer) Render() {
	// Initialize the graphic context on an RGBA image

	renderer.DrawCircle(20, 40) // diameter: 6px, 100x100 = 600x600px + grid 100px = 700x700px
}
