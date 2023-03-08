package main

import (
	"image"
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
)

const TILE_DISPLAY_SIZE = 20

type Renderer struct {
	imageSize  int
	enviroment *Enviroment
	image      *image.RGBA
	context    *draw2dimg.GraphicContext
}

func NewRenderer(enviroment *Enviroment) *Renderer {
	var size = enviroment.size * TILE_DISPLAY_SIZE
	var image = image.NewRGBA(image.Rect(0, 0, size, size))
	var context = draw2dimg.NewGraphicContext(image)

	return &Renderer{
		imageSize:  size,
		enviroment: enviroment,
		image:      image,
		context:    context,
	}
}

func (renderer *Renderer) DrawCircle(x float64, y float64) {
	const halfSize = TILE_DISPLAY_SIZE / 2
	const radius = halfSize - 1
	const circle = 2 * math.Pi

	renderer.context.BeginPath()
	renderer.context.ArcTo(x+halfSize, y+halfSize, radius, radius, 0, circle)
	renderer.context.FillStroke()
}

func (renderer *Renderer) DrawGrid() {
	var size = float64(renderer.imageSize)

	renderer.context.SetStrokeColor(color.RGBA{180, 180, 180, 255})
	renderer.context.SetLineWidth(1)

	// horizontal lines
	for i := 1; i < renderer.enviroment.size; i++ {
		var y = float64(i * TILE_DISPLAY_SIZE)
		renderer.context.MoveTo(0, y)
		renderer.context.LineTo(size, y)
	}

	// vertical lines
	for i := 1; i < renderer.enviroment.size; i++ {
		var x = float64(i * TILE_DISPLAY_SIZE)
		renderer.context.MoveTo(x, 0)
		renderer.context.LineTo(x, size)
	}

	renderer.context.Stroke()
}

func (renderer *Renderer) DrawPopulation(population []*Microbe) {
	renderer.context.SetFillColor(color.RGBA{20, 100, 220, 255})
	renderer.context.SetLineWidth(0)
	for _, microbe := range population {
		var x = float64(microbe.position.x * TILE_DISPLAY_SIZE)
		var y = float64(microbe.position.y * TILE_DISPLAY_SIZE)
		renderer.DrawCircle(x, y)
	}
}

func (renderer *Renderer) RenderScene(population []*Microbe) *image.RGBA {

	renderer.context.SetFillColor(color.White)
	renderer.context.Clear()

	renderer.DrawGrid()
	renderer.DrawPopulation(population)

	return renderer.image // TODO: Return &bytes.Buffer{} instead
}
