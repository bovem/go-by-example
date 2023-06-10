package main

import "fmt"

func sum(x, y int)int{
	return x+y
}

func main(){
	x, y := 1, 2
	fmt.Println(sum(x, y))
}