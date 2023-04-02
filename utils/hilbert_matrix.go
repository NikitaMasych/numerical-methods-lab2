package utils

// HilbertMatrix generates a Hilbert matrix of the desired dimensions
func HilbertMatrix(n int) [][]float64 {
	// Create a 2D slice of floats with n rows and n columns
	mat := make([][]float64, n)
	for i := range mat {
		mat[i] = make([]float64, n)
	}
	// Fill the matrix with Hilbert matrix values
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = 1.0 / float64(i+j+1)
		}
	}
	return mat
}
