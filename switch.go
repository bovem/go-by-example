package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1: 
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	now := time.Now().Weekday()

	switch now {
	case time.Saturday, time.Sunday:
			fmt.Println("Its weekend")
	default:
			fmt.Println("Its a weekday")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("its a boolean")
		case float64:
			fmt.Println("its a float")
		case int:
			fmt.Println("its int")
		default:
			fmt.Println(t, "is something else")
		}
	}

	whatAmI(4.0)
	whatAmI(4)
	whatAmI("Friday")
}