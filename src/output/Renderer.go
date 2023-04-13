package output

import (
	"image"
	"image/color"
	"log"
	"math"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	. "github.com/m-bednar/neuroevolution-thesis/src/enviroment"
	"golang.org/x/image/font/gofont/gomono"
)

const (
	TILE_DISPLAY_SIZE = 18
	FONT_SIZE         = 20
	GRID_LINE_WIDTH   = 1
)

var (
	MICROBE_COLOR    = color.RGBA{30, 120, 240, 255}
	GRID_LINE_COLOR  = color.RGBA{190, 190, 190, 255}
	NONE_TILE_COLOR  = color.White
	SAFE_TILE_COLOR  = color.RGBA{160, 255, 200, 255}
	WALL_TILE_COLOR  = color.RGBA{150, 150, 150, 255}
	SPAWN_TILE_COLOR = color.RGBA{160, 200, 255, 255}
)

type Renderer struct {
	imageSize  int
	font       draw2d.FontData
	enviroment *Enviroment
	background *image.RGBA
}

func LoadGoRegularFont() draw2d.FontData {
	font, err := truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}
	fontData := draw2d.FontData{
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
		return SAFE_TILE_COLOR
	case Wall:
		return WALL_TILE_COLOR
	case Spawn:
		return SPAWN_TILE_COLOR
	default:
		return NONE_TILE_COLOR
	}
}

func GetMicrobeColor(normalized GenomeRepresentation) color.RGBA {
	const alpha = 255
	r, g, b := normalized.red, normalized.green, normalized.blue
	return color.RGBA{r, g, b, alpha}
}

func PredrawTilesOnBackground(context *draw2dimg.GraphicContext, enviroment *Enviroment) {
	for x := 0; x < enviroment.GetSize(); x++ {
		for y := 0; y < enviroment.GetSize(); y++ {
			tile := enviroment.GetTile(NewPosition(x, y))
			color := GetTileColor(tile)
			rx := x * TILE_DISPLAY_SIZE
			ry := y * TILE_DISPLAY_SIZE
			context.SetFillColor(color)
			context.ClearRect(rx, ry, rx+TILE_DISPLAY_SIZE, ry+TILE_DISPLAY_SIZE)
		}
	}
}

func PredrawGridOnBackground(context *draw2dimg.GraphicContext, enviroment *Enviroment, imgSize int) {
	context.SetStrokeColor(GRID_LINE_COLOR)
	context.SetLineWidth(GRID_LINE_WIDTH)

	// horizontal lines
	for i := 1; i < enviroment.GetSize(); i++ {
		y := float64(i * TILE_DISPLAY_SIZE)
		context.MoveTo(0, y)
		context.LineTo(float64(imgSize), y)
	}

	// vertical lines
	for i := 1; i < enviroment.GetSize(); i++ {
		x := float64(i * TILE_DISPLAY_SIZE)
		context.MoveTo(x, 0)
		context.LineTo(x, float64(imgSize))
	}

	context.Stroke()
}

func CreateBackground(enviroment *Enviroment, imgSize int) *image.RGBA {
	background := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))
	context := draw2dimg.NewGraphicContext(background)

	PredrawTilesOnBackground(context, enviroment)
	PredrawGridOnBackground(context, enviroment, imgSize)

	return background
}

func NewRenderer(enviroment *Enviroment) *Renderer {
	size := enviroment.GetSize() * TILE_DISPLAY_SIZE
	background := CreateBackground(enviroment, size)
	font := LoadGoRegularFont()

	return &Renderer{
		imageSize:  size,
		font:       font,
		enviroment: enviroment,
		background: background,
	}
}

func (renderer *Renderer) CreateImageWithContext() (*image.RGBA, *draw2dimg.GraphicContext) {
	img := image.NewRGBA(image.Rect(0, 0, renderer.imageSize, renderer.imageSize))
	context := draw2dimg.NewGraphicContext(img)
	context.SetFontData(renderer.font)
	context.SetFontSize(FONT_SIZE)
	return img, context
}

func DrawCircle(context *draw2dimg.GraphicContext, x float64, y float64) {
	const halfSize = TILE_DISPLAY_SIZE / 2
	const radius = halfSize - GRID_LINE_WIDTH
	const circle = 2 * math.Pi
	context.BeginPath()
	context.ArcTo(x+halfSize, y+halfSize, radius, radius, 0, circle)
	context.FillStroke()
}

func DrawBackground(context *draw2dimg.GraphicContext, img *image.RGBA, background *image.RGBA) {
	// Re-implementation of DrawImage(renderer.background) -
	// does basically the same in this case, but much faster.
	context.Clear()
	copy(img.Pix, background.Pix)
}

func DrawPopulation(context *draw2dimg.GraphicContext, step int, sample CapturedGenerationSample) {
	context.SetLineWidth(0)
	positions := sample.paths[step]

	for i, position := range positions {
		color := GetMicrobeColor(sample.representations[i])
		context.SetFillColor(color)
		x := float64(position.GetX() * TILE_DISPLAY_SIZE)
		y := float64(position.GetY() * TILE_DISPLAY_SIZE)
		DrawCircle(context, x, y)
	}
}

func DrawGenerationNumber(context *draw2dimg.GraphicContext, generation int) {
	context.SetFillColor(color.Black)
	context.SetStrokeColor(color.White)
	context.FillStringAt(strconv.Itoa(generation), 0, FONT_SIZE)
}

func (renderer *Renderer) RenderStep(sample CapturedGenerationSample, step int) *image.RGBA {
	img, context := renderer.CreateImageWithContext()
	DrawBackground(context, img, renderer.background)
	DrawPopulation(context, step, sample)
	DrawGenerationNumber(context, sample.generation)
	return img
}
