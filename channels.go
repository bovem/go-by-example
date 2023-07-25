package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Sends random integer on channel every 2 seconds
func getRandomValue(randomVal chan<- int, numValues int){
	for i:=0;i<numValues;i++{
		value := rand.Int()
		fmt.Println("getRandomValue(): Sent Random Integer", value)

		// Send value to channel
		randomVal <- value

		// 2 seconds interval to ensure that reciever gets enough time
		time.Sleep(time.Second*2)
	}
	close(randomVal)
}

// Listens to channel randomVal and prints received values on terminal
func receiveRandomValue(randomVal <-chan int, numValues int){
	
	// Loop over values in channel
	for value := range randomVal{
		fmt.Println("recieveRandomValue(): Recieved Random Integer", value)
	}
}

func main() {
	randomValueChannel := make(chan int)

	go getRandomValue(randomValueChannel, 5)

	go receiveRandomValue(randomValueChannel, 5)

	// To ensure that main goroutine doesn't end before other
	// two goroutines
	time.Sleep(time.Second*10)
}