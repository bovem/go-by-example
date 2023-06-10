package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Empty list:", a)

	a[2]=2
	a[3]=3
	fmt.Println("New list:", a)

	b := [5]int{3, 3, 234, 2, 2}
	fmt.Println("New list with values:", b)

	var c [5][5]int
	fmt.Println("2D Array:", c)

	fmt.Println("Length of 2D Array:", len(c))

	for i:=0; i<len(c); i++{
		for j:=0; j<len(c); j++{
			c[i][j] = i+j
		}
	}

	fmt.Println("Filled 2D Array:", c)
}