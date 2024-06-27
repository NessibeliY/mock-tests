package main

import "fmt"

type SimpleCalculator struct{}

func (sc SimpleCalculator) Add(a, b int) int {
	return a + b
}

func main() {
	calc := SimpleCalculator{}
	result := calc.Add(2, 3)
	fmt.Println("2 + 3 =", result)
}
