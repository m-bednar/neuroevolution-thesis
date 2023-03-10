package main

import (
	"os"
	"strconv"

	"github.com/wcharczuk/go-chart"
)

const PERCENTAGE_TICKS = 10

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

func MakeChart(collector *DataCollector) {
	var survivability = collector.stats.survivability
	var graph = chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{Top: 20, Left: 20},
		},
		Series: []chart.Series{
			CreateSurvivabilitySeries(survivability),
		},
		XAxis: chart.XAxis{
			Name:      "Generation",
			Style:     chart.StyleShow(),
			NameStyle: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Ticks: CreatePercentageTicks(),
		},
	}

	// Add legend
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	var file, _ = os.Create("output/chart.png")
	var err = graph.Render(chart.PNG, file)

	if err != nil {
		panic(err)
	}

	file.Close()
}
