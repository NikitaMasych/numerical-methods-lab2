package sweep

import (
	"fmt"
	"github.com/NikitaMasych/numerical-methods-lab2/utils"
)

func Example() {
	// Set matrix dimensions and generate a tridiagonal matrix
	n := 5
	matrix := utils.GenerateTridiagonalMatrix(n)

	// Modify the matrix to ensure it is diagonally dominant
	matrix = utils.MakeDiagonallyDominant(matrix)

	// Print the modified matrix
	fmt.Println("Modified Matrix:")
	utils.PrintMatrix(matrix)

	// Solve the system using forward and backward substitution
	solution := solveTridiagonalMatrix(matrix)

	// Print the solution
	fmt.Println("Solution:")
	fmt.Println(solution)
}
