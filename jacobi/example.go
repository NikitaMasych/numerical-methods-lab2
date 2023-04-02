package jacobi

import (
	"fmt"
	"github.com/NikitaMasych/numerical-methods-lab2/utils"
)

func Example() {
	n := 10
	matrix := utils.GenerateHilbertMatrixAndSolution(n)
	matrix = utils.MakeDiagonallyDominant(matrix)
	solution := initializeSolution(n)
	tolerance := 1e-6

	fmt.Println("Matrix:")
	utils.PrintMatrix(matrix)

	fmt.Println("Initial solution:", solution)

	solution = jacobiMethod(matrix, solution, tolerance)

	fmt.Println("Solution:", solution)
}
