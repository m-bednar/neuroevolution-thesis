package neuralnet

import (
	"log"
	"strconv"
	"strings"

	"github.com/myfantasy/mft/im"
)

const (
	NN_OUTPUTS_COUNT = 5
	NN_INPUTS_COUNT  = 6
)

type NeuralNetworkStructure struct {
	layerCount int
	layerWidth int
	maxWidth   int
	layers     []int
	offsets    []int
}

func ParseNeuralNetworkScheme(scheme string) (int, int) {
	const separator = "x"
	var parts = strings.Split(scheme, separator)
	if len(parts) != 2 {
		log.Fatal("Incorrect format of neural network scheme.")
	}
	var count, cerr = strconv.Atoi(parts[0])
	var width, werr = strconv.Atoi(parts[1])
	if cerr != nil || werr != nil {
		log.Fatal("Incorrect format of neural network scheme.")
	}
	return count, width
}

func GetLayerIndexOffset(layers []int, layer int) int {
	var sum = 0
	for i := 0; i < (layer - 1); i++ {
		sum += layers[i] * layers[i+1]
	}
	return sum
}

func ConstructLayersOffsets(layers []int) []int {
	var sum = 0
	var offsets = make([]int, len(layers))
	for i := range layers {
		offsets[i] = sum
		if i+1 < len(layers) {
			sum += layers[i] * layers[i+1]
		}
	}
	return offsets
}

func ConstructLayers(layerCount, layerWidth int) []int {
	var layers = []int{}
	layers = append(layers, NN_INPUTS_COUNT)
	for i := 0; i < layerCount; i++ {
		layers = append(layers, layerWidth)
	}
	layers = append(layers, NN_OUTPUTS_COUNT)
	return layers
}

func NewNeuralNetworkStructure(scheme string) *NeuralNetworkStructure {
	var layerCount, layerWidth = ParseNeuralNetworkScheme(scheme)
	var layers = ConstructLayers(layerCount, layerWidth)
	var offsets = ConstructLayersOffsets(layers)
	return &NeuralNetworkStructure{
		layerCount: layerCount,
		layerWidth: layerWidth,
		maxWidth:   im.MaxS(layers...),
		layers:     layers,
		offsets:    offsets,
	}
}

func (structure *NeuralNetworkStructure) ComputeNumberOfWeights() int {
	var width = structure.layerWidth
	var count = structure.layerCount - 1
	return (NN_INPUTS_COUNT * width) + (NN_OUTPUTS_COUNT * width) + ((width * width) * count)
}

func (structure *NeuralNetworkStructure) GetLayersWidths() []int {
	return structure.layers
}

func (structure *NeuralNetworkStructure) GetMaxLayerWidth() int {
	return structure.maxWidth
}

func (structure *NeuralNetworkStructure) GetLayerOffset(layer int) int {
	return structure.offsets[layer-1]
}
