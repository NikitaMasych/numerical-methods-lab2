package seidel

import (
	"fmt"
	"math"
)

// gaussSeidel solves the linear system Ax = b using the Gauss-Seidel method
func gaussSeidel(a [][]float64, b []float64, x []float64, tol float64, maxIterations int) ([]float64, error) {
	n := len(a)
	for iter := 0; iter < maxIterations; iter++ {
		// Store the previous solution vector
		xPrev := make([]float64, n)
		copy(xPrev, x)

		// Update each component of the solution vector
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += a[i][j] * x[j]
				}
			}
			x[i] = (b[i] - sum) / a[i][i]
		}

		// Check for convergence
		norm := 0.0
		for i := 0; i < n; i++ {
			norm += math.Pow(x[i]-xPrev[i], 2)
		}
		if math.Sqrt(norm) < tol {
			return x, nil
		}
	}
	return nil, fmt.Errorf("gaussSeidel: maximum number of iterations (%d) exceeded", maxIterations)
}
