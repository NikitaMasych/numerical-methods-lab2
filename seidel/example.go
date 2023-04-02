package seidel

import (
	"fmt"
	"github.com/NikitaMasych/numerical-methods-lab2/utils"
)

func Example() {
	n := 3
	a := utils.HilbertMatrix(n)
	fmt.Println("Matrix A:")
	utils.PrintMatrix(a)

	// Modify the matrix to ensure that it is diagonally dominant
	a = utils.MakeDiagonallyDominant(a)
	fmt.Println("Modified Matrix A:")
	utils.PrintMatrix(a)

	// Initialize the solution vector to all zeros
	x := make([]float64, n)

	// Set the maximum number of iterations and the tolerance level
	maxIterations := 1000
	tol := 1e-6

	// Solve the linear system using the Gauss-Seidel method
	x, err := gaussSeidel(a, a[len(a)-1], x, tol, maxIterations)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the solution vector
	fmt.Println("Solution Vector x:")
	utils.PrintVector(x)
}
