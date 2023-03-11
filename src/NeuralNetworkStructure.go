package main

/*
const (
	NN_LAYER_WIDTH  = 8
	NN_NUM_LAYERS   = 2

	NN_NUM_OUTPUTS  = 5
	NN_NUM_INPUTS   = 2
	NN_WEIGHT_LIMIT = 6.0
)
*/

const (
	NN_OUTPUTS_COUNT = 5
	NN_INPUTS_COUNT  = 2
)

type NeuralNetworkStructure struct {
	layerCount int
	layerWidth int
	layers     []int
}

func ConstructLayers(layerCount int, layerWidth int) []int {
	var layers = []int{}
	layers = append(layers, NN_INPUTS_COUNT)
	for i := 0; i < layerCount; i++ {
		layers = append(layers, layerWidth)
	}
	layers = append(layers, NN_OUTPUTS_COUNT)
	return layers
}

func NewNeuralNetworkStructure(layerCount int, layerWidth int) *NeuralNetworkStructure {
	return &NeuralNetworkStructure{
		layerCount: layerCount,
		layerWidth: layerWidth,
		layers:     ConstructLayers(layerCount, layerWidth),
	}
}

func (structure *NeuralNetworkStructure) ComputeNumberOfWeights() int {
	var width = structure.layerWidth
	var count = structure.layerCount - 1
	return (NN_INPUTS_COUNT * width) + (NN_OUTPUTS_COUNT * width) + ((NN_LAYER_WIDTH * width) * count)
}

func (structure *NeuralNetworkStructure) GetLayersWidths() []int {
	return structure.layers
}

func (structure *NeuralNetworkStructure) GetLayerIndexOffset(layer int) int {
	var sum = 0
	for i := 0; i < (layer - 1); i++ {
		sum += structure.layers[i] * structure.layers[i+1]
	}
	return sum
}
