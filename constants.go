package main

import (
	"fmt"
	"math"
)

const s string = "CONSTANT"

func main() {
	fmt.Println("Constant:", s)

	const pi = 3.14

	const d = 3e20 / pi
	fmt.Println("Value of d:", d)
	// fmt.Println("Value of int(d):", int64(d))
	fmt.Println("Value of float64(d):", float64(d))

	d2 := 3e20 / pi
	fmt.Println("Value of d2:", d2)

	fmt.Println("Sine value of pi:", math.Sin(pi))
}