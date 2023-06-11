package main

import "fmt"

type person struct{
	name string
	age int
}

func newPerson(name string) *person{
	var p person
	p.name = name
	return &p
}

func setAge(p *person, age int){
	p.age=age
}

func main(){
	david := newPerson("david")
	fmt.Println(david)

	setAge(david, 35)
	fmt.Println(david)

	dog := struct {
		name string
		age int
	}{
		name:"DoggyDogMcDogFace",
		age:100,
	}

	fmt.Println(dog)
}