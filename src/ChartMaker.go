package main

import (
	"log"
	"os"
	"strconv"

	"github.com/wcharczuk/go-chart"
)

const (
	GENERATION_TICKS = 10
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
	var ticks = make([]chart.Tick, GENERATION_TICKS+1)
	var valueStep = generations / GENERATION_TICKS
	for i := 0; i <= GENERATION_TICKS; i++ {
		ticks[i] = chart.Tick{
			Value: float64(valueStep * i),
			Label: strconv.Itoa(valueStep * i),
		}
	}
	if generations%GENERATION_TICKS != 0 {
		ticks = append(ticks, chart.Tick{
			Value: float64(generations - 1),
			Label: strconv.Itoa(generations - 1),
		})
	}
	return ticks
}

func CreateGraph(averageFitness, highestFitness []float64) chart.Chart {
	var n = len(averageFitness)
	return chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{Top: 20, Left: 20},
		},
		Series: []chart.Series{
			CreateAverageFitnessSeries(averageFitness),
			CreateHighestFitnessSeries(highestFitness),
		},
		XAxis: chart.XAxis{
			Name:      "Generation",
			Style:     chart.StyleShow(),
			NameStyle: chart.StyleShow(),
			Ticks:     CreateGenerationTicks(n),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
	}
}

func (maker *ChartMaker) MakeChart(filename string, collector *DataCollector) {
	var averageFitness = collector.stats.averageFitness
	var highestFitness = collector.stats.highestFitness
	var graph = CreateGraph(averageFitness, highestFitness)

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
