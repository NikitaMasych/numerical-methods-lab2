package sweep

func solveTridiagonalMatrix(matrix [][]float64) []float64 {
	n := len(matrix)

	// Initialize the forward and backward substitution arrays
	a := make([]float64, n)
	b := make([]float64, n)
	c := make([]float64, n)
	d := make([]float64, n)
	x := make([]float64, n)

	// Initialize the forward substitution array
	c[0] = matrix[0][1] / matrix[0][0]
	d[0] = matrix[0][n-1] / matrix[0][0]
	for i := 1; i < n-1; i++ {
		a[i] = matrix[i][i-1]
		b[i] = matrix[i][i]
		c[i] = matrix[i][i+1] / (b[i] - a[i]*c[i-1])
		d[i] = (matrix[i][n-1] - a[i]*d[i-1]) / (b[i] - a[i]*c[i-1])
	}
	a[n-1] = matrix[n-1][n-2]
	b[n-1] = matrix[n-1][n-1]
	d[n-1] = (matrix[n-1][n-1] - a[n-1]*d[n-2]) / (b[n-1] - a[n-1]*c[n-2])

	// Initialize the backward substitution array
	x[n-1] = d[n-1]
	for i := n - 2; i >= 0; i-- {
		x[i] = d[i] - c[i]*x[i+1]
	}

	return x
}
