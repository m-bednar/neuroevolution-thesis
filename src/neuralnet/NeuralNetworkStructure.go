package neuralnet

import (
	"log"
	"strconv"
	"strings"

	"github.com/myfantasy/mft/im"
)

const (
	NN_OUTPUTS_COUNT    = 5
	NN_INPUTS_COUNT     = 6
	NN_SCHEME_SEPARATOR = "x"
)

type NeuralNetworkStructure struct {
	layerCount int
	layerWidth int
	maxWidth   int
	layers     []int
	offsets    []int
}

func ParseNeuralNetworkScheme(scheme string) (int, int) {
	parts := strings.Split(scheme, NN_SCHEME_SEPARATOR)
	if len(parts) != 2 {
		log.Fatal("Incorrect format of neural network scheme.")
	}

	count, cerr := strconv.Atoi(parts[0])
	width, werr := strconv.Atoi(parts[1])
	if cerr != nil || werr != nil {
		log.Fatal("Incorrect format of neural network scheme.")
	}

	return count, width
}

func ConstructLayersOffsets(layers []int) []int {
	sum := 0
	offsets := make([]int, len(layers))
	for i := range layers {
		offsets[i] = sum
		if i+1 < len(layers) {
			sum += layers[i] * layers[i+1]
		}
	}
	return offsets
}

func ConstructLayers(layerCount, layerWidth int) []int {
	layers := []int{}
	layers = append(layers, NN_INPUTS_COUNT)
	for i := 0; i < layerCount; i++ {
		layers = append(layers, layerWidth)
	}
	layers = append(layers, NN_OUTPUTS_COUNT)
	return layers
}

func NewNeuralNetworkStructure(scheme string) *NeuralNetworkStructure {
	count, width := ParseNeuralNetworkScheme(scheme)
	layers := ConstructLayers(count, width)
	offsets := ConstructLayersOffsets(layers)

	return &NeuralNetworkStructure{
		layerCount: count,
		layerWidth: width,
		maxWidth:   im.MaxS(layers...),
		layers:     layers,
		offsets:    offsets,
	}
}

func (structure *NeuralNetworkStructure) ComputeNumberOfWeights() int {
	width := structure.layerWidth
	count := structure.layerCount - 1
	return (NN_INPUTS_COUNT * width) + (NN_OUTPUTS_COUNT * width) + ((width * width) * count)
}

func (structure *NeuralNetworkStructure) GetWeightsIndexSpan(layer, neuron int) (int, int) {
	offset := structure.offsets[layer-1]
	previous := structure.layers[layer-1]
	from := offset + (neuron * previous)
	to := from + previous
	return from, to
}

func (structure *NeuralNetworkStructure) GetLayersWidths() []int {
	return structure.layers
}

func (structure *NeuralNetworkStructure) GetMaxLayerWidth() int {
	return structure.maxWidth
}
