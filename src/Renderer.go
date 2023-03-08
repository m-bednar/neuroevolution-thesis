package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/font/gofont/gomono"
)

const (
	TILE_DISPLAY_SIZE = 20
	FONT_SIZE         = 18
)

var (
	MICROBE_COLOR   = color.RGBA{30, 120, 240, 255}
	GRID_LINE_COLOR = color.RGBA{190, 190, 190, 255}
)

type Renderer struct {
	imageSize  int
	enviroment *Enviroment
	image      *image.RGBA
	context    *draw2dimg.GraphicContext
}

func LoadGoRegularFont() draw2d.FontData {
	var font, err = truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}
	var fontData = draw2d.FontData{
		Name:   "goregular",
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleNormal,
	}
	draw2d.RegisterFont(fontData, font)
	return fontData
}

func NewRenderer(enviroment *Enviroment) *Renderer {
	var size = enviroment.size * TILE_DISPLAY_SIZE
	var image = image.NewRGBA(image.Rect(0, 0, size, size))
	var context = draw2dimg.NewGraphicContext(image)

	context.SetFontData(LoadGoRegularFont())
	context.SetFontSize(FONT_SIZE)

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

	renderer.context.SetStrokeColor(GRID_LINE_COLOR)
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
	renderer.context.SetFillColor(MICROBE_COLOR)
	renderer.context.SetLineWidth(0)
	for _, microbe := range population {
		var x = float64(microbe.position.x * TILE_DISPLAY_SIZE)
		var y = float64(microbe.position.y * TILE_DISPLAY_SIZE)
		renderer.DrawCircle(x, y)
	}
}

func (renderer *Renderer) DrawGenerationNumber(generation int) {
	renderer.context.SetFillColor(color.Black)
	renderer.context.FillStringAt(strconv.Itoa(generation), 0, FONT_SIZE)
}

func (renderer *Renderer) RenderScene(generation int, population []*Microbe) *image.RGBA {
	renderer.context.SetFillColor(color.White)
	renderer.context.Clear()

	renderer.DrawGrid()
	renderer.DrawPopulation(population)
	renderer.DrawGenerationNumber(generation)

	return renderer.image // TODO: Return &bytes.Buffer{} instead
}
