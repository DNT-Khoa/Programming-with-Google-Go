package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Bubble sort
func SortSlice(slice []int, c chan []int) {
	fmt.Println(slice)

	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	c <- slice
}

// Merge function in Merge sort
func MergeTwoSortedArray(slice1, slice2 []int) []int {
	i := 0
	j := 0
	var sortedSubArray []int

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			sortedSubArray = append(sortedSubArray, slice1[i])
			i++
		} else {
			sortedSubArray = append(sortedSubArray, slice2[j])
			j++
		}
	}

	for i < len(slice1) {
		sortedSubArray = append(sortedSubArray, slice1[i])
		i++
	}

	for j < len(slice2) {
		sortedSubArray = append(sortedSubArray, slice2[j])
		j++
	}

	return sortedSubArray
}

func main() {
	c := make(chan []int)
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int

	for {
		fmt.Print("Please enter an integer sequence separated by spaces: ")
		scanner.Scan()
		input := strings.Fields(scanner.Text())

		if len(input) < 4 {
			fmt.Println("Please enter at least 4 integers!")
		} else {
			for _, num := range input {
				numInt, _ := strconv.Atoi(num)
				nums = append(nums, numInt)
			}
			break
		}
	}

	n := len(nums)
	partitionLength := n / 4
	var tempSlice []int
	var unsortedSlice [][]int
	for i := 0; i < 4; i++ {
		if i < 3 {
			tempSlice = nums[partitionLength*i : (partitionLength*i)+partitionLength]
		} else {
			tempSlice = nums[partitionLength*i:]
		}

		go SortSlice(tempSlice, c)
		tempSlice = []int{}
		a := <-c
		unsortedSlice = append(unsortedSlice, a)
	}

	sortedSlice := MergeTwoSortedArray(
		MergeTwoSortedArray(unsortedSlice[0], unsortedSlice[1]),
		MergeTwoSortedArray(unsortedSlice[2], unsortedSlice[3]))
	fmt.Print("Sorted array: ")
	fmt.Print(sortedSlice)
}
