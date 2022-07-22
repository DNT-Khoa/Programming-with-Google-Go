package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var slice []int
	fmt.Print("Please enter a sequence of 10 number separated by spaces: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sequence := scanner.Text()
	numbers := strings.Fields(sequence)
	for _, num := range numbers {
		numInt, _ := strconv.Atoi(num)
		slice = append(slice, numInt)
	}

	BubbleSort(slice)

	for _, num := range slice {
		fmt.Printf("%d ", num)
	}
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}
