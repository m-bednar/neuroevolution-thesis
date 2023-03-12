package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/wcharczuk/go-chart"
)

const (
	PERCENTAGE_TICKS      = 10
	GENERATION_TICKS      = 10
	MOVING_AVG_FRAME_PERC = 0.1
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

func CreateSurvivabilitySeries(data []float64) chart.ContinuousSeries {
	return chart.ContinuousSeries{
		Name:    "Survivability",
		XValues: GenerateContinuousRange(len(data)),
		YValues: data,
	}
}

func CreateMovingAverageSurvivabilitySeries(data []float64) chart.ContinuousSeries {
	return chart.ContinuousSeries{
		Name:    "Survivability [avg]",
		XValues: GenerateContinuousRange(len(data)),
		YValues: GenerateMovingAverage(data),
	}
}

func GenerateMovingAverage(data []float64) []float64 {
	var result = make([]float64, len(data))
	for i := range data {
		result[i] = GetMovingAverageFor(data, i)
	}
	return result
}

func GetMovingAverageFor(data []float64, i int) float64 {
	var dataSize = float64(len(data))
	var frameSize = dataSize * MOVING_AVG_FRAME_PERC
	var half = int(frameSize) / 2
	var start = int(math.Max(float64(i-half), 0))
	var end = int(math.Min(float64(i+half), dataSize))
	var frame = data[start:end]
	var sum = 0.0
	for _, v := range frame {
		sum += v
	}
	return sum / float64(len(frame))
}

func CreatePercentageTicks() []chart.Tick {
	var ticks = make([]chart.Tick, PERCENTAGE_TICKS+1)
	var valueStep = 1.0 / PERCENTAGE_TICKS
	var labelStep = 100 / PERCENTAGE_TICKS
	for i := 0; i <= PERCENTAGE_TICKS; i++ {
		ticks[i] = chart.Tick{
			Value: valueStep * float64(i),
			Label: strconv.Itoa(labelStep*i) + "%",
		}
	}
	return ticks
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

func CreateGraph(survivability []float64) chart.Chart {
	var generations = len(survivability)
	return chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{Top: 20, Left: 20},
		},
		Series: []chart.Series{
			CreateSurvivabilitySeries(survivability),
			CreateMovingAverageSurvivabilitySeries(survivability),
		},
		XAxis: chart.XAxis{
			Name:      "Generation",
			Style:     chart.StyleShow(),
			NameStyle: chart.StyleShow(),
			Ticks:     CreateGenerationTicks(generations),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Ticks: CreatePercentageTicks(),
		},
	}
}

func (maker *ChartMaker) MakeChart(filename string, collector *DataCollector) {
	var survivability = collector.stats.survivability
	var graph = CreateGraph(survivability)

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
