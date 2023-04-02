package utils

import (
	"math/rand"
	"time"
)

// GenerateVector generates a random vector with the desired dimensions.
func GenerateVector(n int) []float64 {
	// create empty vector
	b := make([]float64, n)

	// fill vector with random values
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		b[i] = rand.Float64()
	}

	return b
}
