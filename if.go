package main

import "fmt"

func main() {
	i := 1

	if i<2 {
		fmt.Println("i is less than 2")
	}
	
	i = 2

	if i<2 {
		fmt.Println("i is less than 2")
	} else if i==2 {
		fmt.Println("i is 2")
	} else {
		fmt.Println("i is greater than 2")
	}
}