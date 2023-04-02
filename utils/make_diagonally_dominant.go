package utils

import "math/rand"

func MakeDiagonallyDominant(matrix [][]float64) [][]float64 {
	// Loop through each row of the matrix
	for i, row := range matrix {
		// Calculate the sum of the absolute values of the off-diagonal elements
		offDiagonalSum := 0.0
		for j, elem := range row {
			if j != i {
				offDiagonalSum += elem
			}
		}

		// Modify the diagonal element to ensure diagonal dominance
		if matrix[i][i] < offDiagonalSum {
			matrix[i][i] = offDiagonalSum + rand.Float64()
		}
	}

	return matrix
}
