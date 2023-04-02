package gauss

// LUPSolve solves the system L*U*x = P*b using forward and backward substitution.
// L and U are the lower and upper triangular matrices from the LU decomposition with partial pivoting,
// P is the permutation matrix, and b is the vector of constants.
// Returns the solution vector x.
func LUPSolve(L, U, P [][]float64, b []float64) []float64 {
	// get dimensions
	n := len(L)

	// initialize x
	x := make([]float64, n)

	// solve L*y = P*b using forward substitution
	y := make([]float64, n)
	pb := make([]float64, n)
	for i := 0; i < n; i++ {
		pb[i] = b[int(P[i][0])]
	}
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < i; j++ {
			sum += L[i][j] * y[j]
		}
		y[i] = pb[i] - sum
	}

	// solve U*x = y using backward substitution
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += U[i][j] * x[j]
		}
		x[i] = (y[i] - sum) / U[i][i]
	}

	return x
}
