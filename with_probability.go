package random

import (
	"math/rand"
)

// GetRandomWithProbabilities returns random item
// from slice of any data type with given probalities
func GetRandomWithProbabilities(items []any, probabilities []float64) any {
	if len(items) == 0 || len(probabilities) == 0 || len(items) != len(probabilities) {
		return nil
	}

	var sumProbabilities float64
	for _, value := range probabilities {
		if value < 0 {
			return nil
		}
		sumProbabilities += value
	}

	if sumProbabilities == 0 {
		return nil
	}

	randValue := randomFloat64(sumProbabilities)
	accumulated := 0.0

	for i, item := range items {
		accumulated += probabilities[i]
		if randValue < accumulated {
			return item
		}
	}

	return items[len(items)-1]
}

// GetRandomStructWithProbabilities returns random item
// from slice of structures with given probabilities
func GetRandomStructWithProbabilities(items []interface{ GetProbability() float64 }) any {
	if len(items) == 0 {
		return nil
	}

	var sumProbabilities float64
	for _, v := range items {
		prob := v.GetProbability()
		if prob < 0 {
			return nil
		}
		sumProbabilities += prob
	}

	if sumProbabilities == 0 {
		return nil
	}

	randValue := randomFloat64(sumProbabilities)
	accumulated := 0.0

	for _, item := range items {
		accumulated += item.GetProbability()
		if randValue < accumulated {
			return item
		}
	}

	return items[len(items)-1]
}

// GetRandomMapItemWithProbabilities returns random item
// from a map where values are probabilities
func GetRandomMapItemWithProbabilities(items map[string]float64) string {
	if len(items) == 0 {
		return ""
	}

	var sumProbabilities float64
	for _, v := range items {
		if v < 0 {
			return ""
		}
		sumProbabilities += v
	}

	if sumProbabilities == 0 {
		return ""
	}

	randValue := randomFloat64(sumProbabilities)
	accumulated := 0.0

	var lastKey string
	for k, v := range items {
		lastKey = k
		accumulated += v
		if randValue < accumulated {
			return k
		}
	}

	return lastKey
}

// GetRandomMapItemWithPercent returns random item
// from a map where values are drop percentages
func GetRandomMapItemWithPercent(items map[string]float64) string {
	if len(items) == 0 {
		return ""
	}

	const maxAttempts = 100

	for i := 0; i < maxAttempts; i++ {
		for k, v := range items {
			if randomFloat64(100) < v {
				return k
			}
		}
	}

	for k := range items {
		return k
	}

	return ""
}

// Random float64 number in given max
func randomFloat64(max float64) float64 {
	return rand.Float64() * max
}
