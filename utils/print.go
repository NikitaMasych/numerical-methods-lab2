package utils

import "fmt"

// PrintMatrix prints the given matrix in a human-readable format.
func PrintMatrix(A [][]float64) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			fmt.Printf("%10.4f ", A[i][j])
		}
		fmt.Println()
	}
}

// PrintVector prints the given vector in a human-readable format.
func PrintVector(v []float64) {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%10.4f ", v[i])
	}
	fmt.Println()
}
