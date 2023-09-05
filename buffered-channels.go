package main

import "fmt"

func main() {
	exampleBufferedChan := make(chan string, 3)

	// Sending values to buffered channel
	exampleBufferedChan<-"c"
	exampleBufferedChan<-"b"
	exampleBufferedChan<-"a"

	// Recieving values from buffered channel
	fmt.Println("Value 1 from exampleBufferedChan:",<-exampleBufferedChan)
	fmt.Println("Value 2 from exampleBufferedChan:",<-exampleBufferedChan)
	fmt.Println("Value 3 from exampleBufferedChan:",<-exampleBufferedChan)
}