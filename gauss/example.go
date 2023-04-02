package gauss

import (
	"fmt"
	"github.com/NikitaMasych/numerical-methods-lab2/utils"
)

func Example() {
	// set dimensions
	n := 3

	// generate random or Hilbert matrix
	A := utils.HilbertMatrix(n)

	// perform LU decomposition with partial pivoting
	L, U, P := LUPDecompose(A)

	// generate random vector
	b := utils.GenerateVector(n)

	// solve the system
	x := LUPSolve(L, U, P, b)

	// print results
	fmt.Println("A:")
	utils.PrintMatrix(A)
	fmt.Println("L:")
	utils.PrintMatrix(L)
	fmt.Println("U:")
	utils.PrintMatrix(U)
	fmt.Println("P:")
	utils.PrintMatrix(P)
	fmt.Println("b:")
	utils.PrintVector(b)
	fmt.Println("x:")
	utils.PrintVector(x)
}
