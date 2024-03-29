/**
 * @project neuroevolution/utils
 * @file    SliceUtils.go
 * @author  Martin Bednář (xbedna77@fit.vut.cz)
 */

package utils

import (
	"math"
)

func SliceMaxValue(values []float64) float64 {
	max := -math.MaxFloat64
	for _, value := range values {
		max = math.Max(max, value)
	}
	return max
}

func SliceMinValue(values []float64) float64 {
	min := math.MaxFloat64
	for _, value := range values {
		min = math.Min(min, value)
	}
	return min
}

func Clone[T any](items []T) []T {
	new := make([]T, len(items))
	copy(new, items)
	return new
}
