package gauss

import "math"

// LUPDecompose performs LU decomposition with partial pivoting on a matrix A.
// It returns the lower triangular matrix L, the upper triangular matrix U,
// and the permutation matrix P such that P*A = L*U.
func LUPDecompose(A [][]float64) ([][]float64, [][]float64, [][]float64) {
	// get dimensions
	n := len(A)

	// initialize L, U, and P
	L := make([][]float64, n)
	U := make([][]float64, n)
	P := make([][]float64, n)
	for i := range L {
		L[i] = make([]float64, n)
		U[i] = make([]float64, n)
		P[i] = make([]float64, n)
		P[i][i] = 1
	}

	// copy A to U
	for i := range A {
		for j := range A[i] {
			U[i][j] = A[i][j]
		}
	}

	// perform partial pivoting
	for k := 0; k < n-1; k++ {
		// find pivot row
		max := 0.0
		imax := k
		for i := k; i < n; i++ {
			if math.Abs(U[i][k]) > max {
				max = math.Abs(U[i][k])
				imax = i
			}
		}

		// swap rows
		if imax != k {
			U[k], U[imax] = U[imax], U[k]
			P[k], P[imax] = P[imax], P[k]
			if k > 0 {
				for i := 0; i < k; i++ {
					L[k][i], L[imax][i] = L[imax][i], L[k][i]
				}
			}
		}

		// eliminate
		for i := k + 1; i < n; i++ {
			L[i][k] = U[i][k] / U[k][k]
			for j := k; j < n; j++ {
				U[i][j] -= L[i][k] * U[k][j]
			}
		}
	}

	// fill diagonal of L with 1's
	for i := range L {
		L[i][i] = 1
	}

	return L, U, P
}
