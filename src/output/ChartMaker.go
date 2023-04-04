package output

import (
	"log"
	"os"
	"strconv"

	"github.com/myfantasy/mft/im"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

const (
	GENERATION_TICKS_COUNT  = 10
	VALUE_TICKS_COUNT       = 5
	VALUE_TICK_PRECISION    = 2
	WINDOW_FILTER_SIZE_PERC = 0.01
)

type ChartMaker struct {
	collector *DataCollector
}

func NewChartMaker(collector *DataCollector) *ChartMaker {
	return &ChartMaker{
		collector: collector,
	}
}

func GenerateContinuousRange(x int) []float64 {
	r := make([]float64, x)
	for i := 0; i < x; i++ {
		r[i] = float64(i)
	}
	return r
}

func CreateAverageFitnessSeries(data []float64) chart.ContinuousSeries {
	return chart.ContinuousSeries{
		Name:    "Fitness [avg.]",
		XValues: GenerateContinuousRange(len(data)),
		YValues: data,
	}
}

func CreateHighestFitnessSeries(data []float64) chart.ContinuousSeries {
	return chart.ContinuousSeries{
		Name:    "Fitness [high.]",
		XValues: GenerateContinuousRange(len(data)),
		YValues: data,
	}
}

func CreateGenerationTicks(generations int) []chart.Tick {
	ticks := make([]chart.Tick, GENERATION_TICKS_COUNT+1)
	step := generations / GENERATION_TICKS_COUNT
	for i := 0; i <= GENERATION_TICKS_COUNT; i++ {
		ticks[i] = chart.Tick{
			Value: float64(step * i),
			Label: strconv.Itoa(step * i),
		}
	}
	if generations%GENERATION_TICKS_COUNT != 0 {
		ticks = append(ticks, chart.Tick{
			Value: float64(generations - 1),
			Label: strconv.Itoa(generations - 1),
		})
	}
	return ticks
}

func CreateValueTicks() []chart.Tick {
	ticks := make([]chart.Tick, VALUE_TICKS_COUNT+1)
	step := 1.0 / VALUE_TICKS_COUNT
	for i := 0; i <= VALUE_TICKS_COUNT; i++ {
		value := step * float64(i)
		ticks[i] = chart.Tick{
			Value: value,
			Label: strconv.FormatFloat(value, 'f', VALUE_TICK_PRECISION, 64),
		}
	}
	return ticks
}

func WindowFilter(data []float64, index int) float64 {
	size := int(float64(len(data)) * WINDOW_FILTER_SIZE_PERC)
	half := size / 2
	from := im.Max(0, index-half)
	to := im.Min(len(data), index+half)

	window := data[from:to]
	sum := 0.0
	for _, v := range window {
		sum += v
	}

	return sum / float64(len(window))
}

func WindowFilterAll(data []float64) []float64 {
	filtered := make([]float64, len(data))
	for i := range data {
		filtered[i] = WindowFilter(data, i)
	}
	return filtered
}

func CreateGraph(averageFitnesses, highestFitnesses []float64) *chart.Chart {
	n := len(averageFitnesses)
	return &chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{Top: 20, Left: 20},
		},
		Series: []chart.Series{
			CreateAverageFitnessSeries(WindowFilterAll(averageFitnesses)),
			CreateHighestFitnessSeries(WindowFilterAll(highestFitnesses)),
		},
		XAxis: chart.XAxis{
			Name:      "Generation",
			Style:     chart.StyleShow(),
			NameStyle: chart.StyleShow(),
			Ticks:     CreateGenerationTicks(n),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Ticks: CreateValueTicks(),
		},
	}
}

func AddLegendToGraph(graph *chart.Chart) {
	gray := drawing.Color{R: 140, G: 140, B: 140, A: 255}
	padding := chart.Box{Top: 1, Left: 1}
	legend := chart.LegendLeft(graph, chart.Style{
		FontSize:    9,
		StrokeColor: gray,
		Padding:     padding,
	})
	graph.Elements = []chart.Renderable{legend}
}

func SaveGraphToFile(graph *chart.Chart, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := graph.Render(chart.PNG, file); err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func (maker *ChartMaker) MakeChart(filename string) {
	averageFitnesses := maker.collector.GetAverageFitnesses()
	highestFitnesses := maker.collector.GetHighestFitnesses()
	graph := CreateGraph(averageFitnesses, highestFitnesses)
	AddLegendToGraph(graph)
	SaveGraphToFile(graph, filename)
}
