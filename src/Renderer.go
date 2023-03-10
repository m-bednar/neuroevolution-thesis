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
	TILE_DISPLAY_SIZE = 12
	FONT_SIZE         = 18
)

var (
	MICROBE_COLOR   = color.RGBA{30, 120, 240, 255}
	GRID_LINE_COLOR = color.RGBA{190, 190, 190, 255}
	NONE_TILE_COLOR = color.White
	SAFE_FILE_COLOR = color.RGBA{190, 255, 210, 255}
)

type Renderer struct {
	imageSize  int
	enviroment *Enviroment
	image      *image.RGBA
	background *image.RGBA
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
		Style:  draw2d.FontStyleBold,
	}
	draw2d.RegisterFont(fontData, font)
	return fontData
}

func GetTileColor(tile TileType) color.Color {
	switch tile {
	case Safe:
		return SAFE_FILE_COLOR
	default:
		return NONE_TILE_COLOR
	}
}

func PredrawTilesOnBackground(context *draw2dimg.GraphicContext, enviroment *Enviroment) {
	for x := 0; x < enviroment.size; x++ {
		for y := 0; y < enviroment.size; y++ {
			var tile = enviroment.GetTile(NewPosition(x, y))
			var color = GetTileColor(tile)
			var rx = x * TILE_DISPLAY_SIZE
			var ry = y * TILE_DISPLAY_SIZE
			context.SetFillColor(color)
			context.ClearRect(rx, ry, rx+TILE_DISPLAY_SIZE, ry+TILE_DISPLAY_SIZE)
		}
	}
}

func PredrawGridOnBackground(context *draw2dimg.GraphicContext, enviroment *Enviroment, imgSize int) {
	context.SetStrokeColor(GRID_LINE_COLOR)
	context.SetLineWidth(1)

	// horizontal lines
	for i := 1; i < enviroment.size; i++ {
		var y = float64(i * TILE_DISPLAY_SIZE)
		context.MoveTo(0, y)
		context.LineTo(float64(imgSize), y)
	}

	// vertical lines
	for i := 1; i < enviroment.size; i++ {
		var x = float64(i * TILE_DISPLAY_SIZE)
		context.MoveTo(x, 0)
		context.LineTo(x, float64(imgSize))
	}

	context.Stroke()
}

func CreateBackground(enviroment *Enviroment, imgSize int) *image.RGBA {
	var background = image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))
	var context = draw2dimg.NewGraphicContext(background)

	PredrawTilesOnBackground(context, enviroment)
	PredrawGridOnBackground(context, enviroment, imgSize)

	return background
}

func NewRenderer(enviroment *Enviroment) *Renderer {
	var size = enviroment.size * TILE_DISPLAY_SIZE
	var img = image.NewRGBA(image.Rect(0, 0, size, size))
	var context = draw2dimg.NewGraphicContext(img)
	var background = CreateBackground(enviroment, size)

	context.SetFontData(LoadGoRegularFont())
	context.SetFontSize(FONT_SIZE)

	return &Renderer{
		imageSize:  size,
		enviroment: enviroment,
		image:      img,
		background: background,
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

func (renderer *Renderer) DrawBackground() {
	// Own re-implementation of DrawImage(renderer.background)
	// Does the basically same in this case, but much faster.
	renderer.context.Clear()
	copy(renderer.image.Pix, renderer.background.Pix)
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
	renderer.DrawBackground()
	renderer.DrawPopulation(population)
	renderer.DrawGenerationNumber(generation)
	return renderer.image
}
