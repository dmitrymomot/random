package random

import (
	"math/rand"
	"time"
)

// GetRandomWithProbabilities returns random item from slice with given probalities
func GetRandomWithProbabilities(items []interface{}, probabilities []float64) interface{} {
	var (
		probability float64
		item        interface{}
		i           int

		sum = sumSlice(probabilities)
	)

	for i, item = range items {
		probability = probabilities[i]

		if randomFloat64(sum) <= probability {
			break
		}
	}

	return item
}

// Random float64 number in given max
func randomFloat64(max float64) float64 {
	rand.Seed(time.Now().Unix())

	return rand.Float64() * max
}

// sum of all values in given slice
func sumSlice(values []float64) float64 {
	var result float64

	for _, value := range values {
		result += value
	}

	return result
}
