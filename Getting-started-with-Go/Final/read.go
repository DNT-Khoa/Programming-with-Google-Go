package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	type Person struct {
		fName string
		lName string
	}

	var mySlice []Person
	var fileName string
	fmt.Println("Please enter the file name: ")
	fmt.Scan(&fileName)

	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		mySlice = append(mySlice, Person{fName: words[0], lName: words[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, p := range mySlice {
		fmt.Printf("%s %s\n", p.fName, p.lName)
	}
}
