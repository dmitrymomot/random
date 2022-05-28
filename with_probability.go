package random

import (
	"math/rand"
	"time"
)

// GetRandomWithProbabilities returns random item
// from slice of any data type with given probalities
func GetRandomWithProbabilities(items []interface{}, probabilities []float64) interface{} {
	var (
		probability      float64
		sumProbabilities float64
		item             interface{}
		i                int
	)

	for _, value := range probabilities {
		sumProbabilities += value
	}

	for i, item = range items {
		probability = probabilities[i]

		if randomFloat64(sumProbabilities) <= probability {
			break
		}
	}

	return item
}

// GetRandomStructWithProbabilities returns random item
// from slice of structures with given probabilities
func GetRandomStructWithProbabilities(items []interface{ GetProbability() float64 }) interface{} {
	var (
		sumProbabilities float64
		result           interface{}
	)

	for _, v := range items {
		sumProbabilities += v.GetProbability()
	}

	for _, item := range items {
		result = item
		if randomFloat64(sumProbabilities) <= item.GetProbability() {
			break
		}
	}

	return result
}

// GetRandomMapItemWithProbabilities returns random item
// from a map where values are probabilities
func GetRandomMapItemWithProbabilities(items map[string]float64) string {
	var (
		sumProbabilities float64
		counter          float64
		result           string
	)

	for _, v := range items {
		sumProbabilities += v
	}

	for k, v := range items {
		counter += v
		if randomFloat64(sumProbabilities) < counter {
			result = k
			break
		}
	}

	if result == "" {
		result = GetRandomMapItemWithProbabilities(items)
	}

	return result
}

// GetRandomMapItemWithProbabilitiesOf100Percent returns random item
// from a map where values are probabilities
func GetRandomMapItemWithProbabilitiesOf100Percent(items map[string]float64) string {
	var (
		sumProbabilities float64 = 100.0
		counter          float64
		result           string
	)

	for k, v := range items {
		counter += v
		if randomFloat64(sumProbabilities) < counter {
			result = k
			break
		}
	}

	if result == "" {
		result = GetRandomMapItemWithProbabilitiesOf100Percent(items)
	}

	return result
}

// Random float64 number in given max
func randomFloat64(max float64) float64 {
	rand.Seed(time.Now().UnixNano())

	return rand.Float64() * max
}
