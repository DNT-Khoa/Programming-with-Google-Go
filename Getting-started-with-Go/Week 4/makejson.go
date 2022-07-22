package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Please enter an address: ")
	scanner.Scan()
	address := scanner.Text()

	myObject := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonObject, _ := json.Marshal(myObject)

	fmt.Println(string(jsonObject))
}
