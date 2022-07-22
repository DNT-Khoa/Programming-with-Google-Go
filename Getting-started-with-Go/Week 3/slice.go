package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	for {
		var input string
		fmt.Print("Enter an integer: ")
		fmt.Scan(&input)

		if input == "X" {
			break
		} else {
			num, _ := strconv.Atoi(input)
			slice = append(slice, num)
			sort.Ints(slice)
			fmt.Println(slice)
		}
	}
}
