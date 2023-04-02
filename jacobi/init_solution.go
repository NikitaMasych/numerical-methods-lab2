package jacobi

import "math/rand"

func initializeSolution(n int) []float64 {
	solution := make([]float64, n)
	for i := 0; i < n; i++ {
		solution[i] = rand.Float64()
	}
	return solution
}
