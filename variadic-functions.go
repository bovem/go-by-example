package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, value := range nums{
		sum += value
	}
	return sum
}

func main() {
	arr1 := []int{11, 12, 23, 21, 12}
	fmt.Println("Sum of array:", arr1, "is:", sum(arr1...))
}