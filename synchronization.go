package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomInt(doneFlag chan bool){
	fmt.Println("Goroutine for getRandomInt() started")
	fmt.Println("Some random integer: ", rand.Int())
	time.Sleep(time.Second*5)

	// Will send value to the channel
	// once the function is finished
	doneFlag<-true
}

func main(){
	done := make(chan bool)

	go getRandomInt(done)

	// Main Goroutine will wait for value on the channel
	<-done

	fmt.Println("Main Goroutine has ended")
}

// Output
// Goroutine for getRandomInt() started
// Some random integer:  4571662314887367842
// Main Goroutine has ended