package main

import (
	"fmt"
	"time"
)

var sharedInt = 0

func IncrementValue() {
	fmt.Println("Increment by 1")
	sharedInt += 1
	fmt.Printf("Value has been incremented: %d\n", sharedInt)
}

func DoubleValue() {
	fmt.Println("Double value")
	sharedInt *= 2
	fmt.Printf("Value has been doubled: %d\n", sharedInt)
}

func PrintValue() {
	fmt.Printf("Value of sharedInt: %d\n", sharedInt)
}

/*
By executing: go run -race racecondition.go, you will get a data race warning
*/
func main() {
	go IncrementValue()
	go DoubleValue()
	go PrintValue()
	time.Sleep(5 * time.Second)
}

/*
1. What is a race condition and how it can occur ?
A race condition is an undesirable condition that happens when multiple threads or go routines trying to access and modify the same data.
In other words, a race condition can happen when there is a communication between different concurrent tasks
*/
