package utils

import (
	"math/rand"
	"time"
)

// RandomMatrix generates a random matrix of the desired dimensions
func RandomMatrix(n int) [][]float64 {
	// Create a 2D slice of floats with n rows and n columns
	mat := make([][]float64, n)
	for i := range mat {
		mat[i] = make([]float64, n)
	}
	// Fill the matrix with random values between 0 and 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = rand.Float64()
		}
	}
	return mat
}
