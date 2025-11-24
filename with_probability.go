package random

import (
	"math/rand"
)

// GetRandomWithProbabilities returns a random item from a slice with given probabilities.
// Probabilities are relative weights and do not need to sum to any specific value.
// Returns the zero value of T if inputs are invalid (empty, mismatched lengths, negative probabilities).
func GetRandomWithProbabilities[T any](items []T, probabilities []float64) T {
	var zero T
	if len(items) == 0 || len(probabilities) == 0 || len(items) != len(probabilities) {
		return zero
	}

	var sumProbabilities float64
	for _, value := range probabilities {
		if value < 0 {
			return zero
		}
		sumProbabilities += value
	}

	if sumProbabilities == 0 {
		return zero
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

// GetRandomStructWithProbabilities returns a random item from a slice of structures
// that implement the GetProbability() float64 method.
// Probabilities are relative weights and do not need to sum to any specific value.
// Returns the zero value of T if inputs are invalid (empty, negative probabilities).
func GetRandomStructWithProbabilities[T interface{ GetProbability() float64 }](items []T) T {
	var zero T
	if len(items) == 0 {
		return zero
	}

	var sumProbabilities float64
	for _, v := range items {
		prob := v.GetProbability()
		if prob < 0 {
			return zero
		}
		sumProbabilities += prob
	}

	if sumProbabilities == 0 {
		return zero
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

// GetRandomMapItemWithPercent returns a random key from a map using weighted selection.
// The values are treated as relative weights (percentages) and do not need to sum to 100.
// This uses proper weighted selection, guaranteeing one item is always selected based on
// relative probabilities, making it suitable for lootbox mechanics.
// Returns empty string if the map is empty or contains only negative values.
func GetRandomMapItemWithPercent(items map[string]float64) string {
	if len(items) == 0 {
		return ""
	}

	// Calculate sum and validate
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

	// Weighted random selection
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

// randomFloat64 returns a random float64 value in the range [0, max).
func randomFloat64(max float64) float64 {
	return rand.Float64() * max
}
