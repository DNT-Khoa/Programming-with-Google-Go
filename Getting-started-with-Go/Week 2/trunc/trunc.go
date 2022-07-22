package main

import "fmt"

func main() {
	var fInput float64
	fmt.Println("Please enter a floating point number ?")
	fmt.Scan(&fInput)
	fmt.Println(int(fInput))
}
