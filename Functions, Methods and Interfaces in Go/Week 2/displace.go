package main

import (
	"fmt"
	"time"
)

func GenDisplayFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		s := (1.0/2)*a*(t*t) + v0*t + s0
		time.Sleep(time.Duration(t * float64(time.Second)))
		return s
	}
}

func main() {
	var a float64
	var v0 float64
	var s0 float64
	fmt.Print("Please enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Please enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Please enter initial displacement: ")
	fmt.Scan(&s0)

	myFunc := GenDisplayFn(a, v0, s0)

	var t float64
	fmt.Print("Please enter a time (in seconds): ")
	fmt.Scan(&t)
	fmt.Printf("Result: %f", myFunc(t))
}
