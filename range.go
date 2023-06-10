package main

import "fmt"

func main() {
	n := [5]int{12, 12, 123, 213, 213}
	fmt.Println("Array:", n)

	for i, val := range n {
		fmt.Println("Index:", i, " Value:", val)
	}

	m := map[string]int{"A":1, "B":2, "C":3}
	fmt.Println("HashMap:", m)

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	for index, unicode := range "Hello, World"{
		fmt.Println("Index:", index, "Unicode:", unicode)
	}
}