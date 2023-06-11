package main

import "fmt"

func zeroConverter(val int){
	val=0
}

func zeroConverterPtr(val *int){
	*val=0
}

func main() {
	a := 5
	fmt.Println("Value of a:", a)

	zeroConverter(a)
	fmt.Println("Value of a after zeroConverter", a)

	zeroConverterPtr(&a)
	fmt.Println("Value of a after zeroConverterPtr", a)
}