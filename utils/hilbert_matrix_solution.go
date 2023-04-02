package utils

import "math/rand"

// GenerateHilbertMatrixAndSolution generates a Hilbert matrix of the desired dimensions.
// Initializes solution column with n index
func GenerateHilbertMatrixAndSolution(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			matrix[i][j] = 1.0 / float64(i+j+1)
		}
		matrix[i][n] = rand.Float64()
	}
	return matrix
}
