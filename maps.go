package main

import "fmt"

func main(){
	m := make(map[string]int)
	
	m["a"]=1
	m["b"]=2

	fmt.Println("Map:", m)

	value1 := m["a"]
	value2 := m["b"]

	fmt.Println("Values:", value1, value2)
	fmt.Println("Length of the Map:", len(m))

	delete(m, "b")
	fmt.Println("Map after deletion:", m)

	v1, v2 := m["b"]
	fmt.Println("Value of v1:", v1)
	fmt.Println("Value of v2:", v2)

	newMap := map[string]int{"a":1, "b":2}
	fmt.Println("New Map:", newMap)
}