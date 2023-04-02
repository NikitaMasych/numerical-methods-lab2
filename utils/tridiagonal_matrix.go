package utils

import "math/rand"

func GenerateTridiagonalMatrix(n int) [][]float64 {
	// Initialize a n x n matrix with all zeroes
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}

	// Fill in the tridiagonal elements randomly
	for i := 0; i < n; i++ {
		matrix[i][i] = rand.Float64() * 10
		if i > 0 {
			matrix[i][i-1] = rand.Float64() * 10
		}
		if i < n-1 {
			matrix[i][i+1] = rand.Float64() * 10
		}
	}

	return matrix
}
