package main

import "fmt"

func sumAndDiff(x, y int) (int, int){
	return x+y, x-y
}

func main(){
	x, y := 1, 2
	sum, diff := sumAndDiff(x, y)
	fmt.Println("Sum and Difference:", sum, ",", diff)
}