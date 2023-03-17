package main

import (
	"log"
	"os"
	"strconv"

	"github.com/myfantasy/mft/im"
	"github.com/wcharczuk/go-chart"
)

const (
	GENERATION_TICKS_COUNT  = 10
	VALUE_TICKS_COUNT       = 10
	VALUE_TICK_PRECISION    = 2
	WINDOW_FILTER_SIZE_PERC = 0.02
)

type ChartMaker struct{}

func NewChartMaker() *ChartMaker {
	return &ChartMaker{}
}

func GenerateContinuousRange(x int) []float64 {
	var r = make([]float64, x)
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
	var ticks = make([]chart.Tick, GENERATION_TICKS_COUNT+1)
	var step = generations / GENERATION_TICKS_COUNT
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
	var ticks = make([]chart.Tick, VALUE_TICKS_COUNT+1)
	var step = 1.0 / VALUE_TICKS_COUNT
	for i := 0; i <= VALUE_TICKS_COUNT; i++ {
		var value = step * float64(i)
		ticks[i] = chart.Tick{
			Value: value,
			Label: strconv.FormatFloat(value, 'f', VALUE_TICK_PRECISION, 64),
		}
	}
	return ticks
}

func WindowFilter(data []float64, index int) float64 {
	var size = int(float64(len(data)) * WINDOW_FILTER_SIZE_PERC)
	var half = size / 2
	var from = im.Max(0, index-half)
	var to = im.Min(len(data), index+half)

	var window = data[from:to]
	var sum = 0.0
	for _, v := range window {
		sum += v
	}

	return sum / float64(len(window))
}

func WindowFilterAll(data []float64) []float64 {
	var filtered = make([]float64, len(data))
	for i := range data {
		filtered[i] = WindowFilter(data, i)
	}
	return filtered
}

func CreateGraph(averageFitnesses, highestFitnesses []float64) chart.Chart {
	var n = len(averageFitnesses)
	return chart.Chart{
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

func (maker *ChartMaker) MakeChart(filename string, collector *DataCollector) {
	var averageFitnesses = collector.GetAverageFitnesses()
	var highestFitnesses = collector.GetHighestFitnesses()
	var graph = CreateGraph(averageFitnesses, highestFitnesses)

	// Add legend to chart
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	var file, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := graph.Render(chart.PNG, file); err != nil {
		log.Fatal(err)
	}

	file.Close()
}
