package jacobi

import "math"

func jacobiMethod(matrix [][]float64, solution []float64, tolerance float64) []float64 {
	n := len(matrix)
	newSolution := make([]float64, n)
	for i := 0; i < n; i++ {
		newSolution[i] = solution[i]
	}

	diff := tolerance + 1.0
	for diff > tolerance {
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j {
					sum += matrix[i][j] * solution[j]
				}
			}
			newSolution[i] = (matrix[i][n] - sum) / matrix[i][i]
		}

		diff = 0.0
		for i := 0; i < n; i++ {
			d := newSolution[i] - solution[i]
			diff += d * d
		}
		diff = math.Sqrt(diff)

		for i := 0; i < n; i++ {
			solution[i] = newSolution[i]
		}
	}

	return solution
}
