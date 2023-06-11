package main

import "fmt"

type quad struct{
	height int
	width int
}

func (q *quad) area() int{
	return q.height*q.width
}

func (q quad) perimeter() int{
	return (2*q.height)+(2*q.width)
}

func main(){
	r := quad{height:4, width:5}
	fmt.Println("Area of quadrilateral:", r.area())
	fmt.Println("Perimeter of quadrilateral:", r.perimeter())
}