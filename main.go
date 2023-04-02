package main

import (
	"fmt"
	"github.com/NikitaMasych/numerical-methods-lab2/gauss"
	"github.com/NikitaMasych/numerical-methods-lab2/jacobi"
	"github.com/NikitaMasych/numerical-methods-lab2/seidel"
	"github.com/NikitaMasych/numerical-methods-lab2/sweep"
	"strings"
)

const (
	separator = "-"
	amount    = 60
)

func main() {
	delimiter := strings.Repeat(separator, amount)
	fmt.Println(delimiter, "\nGauss:")
	gauss.Example()
	fmt.Println(delimiter, "\nSweep:")
	sweep.Example()
	fmt.Println(delimiter, "\nJacobi:")
	jacobi.Example()
	fmt.Println(delimiter, "\nSeidel:")
	seidel.Example()
}
