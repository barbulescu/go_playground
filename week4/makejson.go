package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter the name: ")
	var name string
	fmt.Fscan(reader, &name)

	fmt.Print("Please enter the address: ")
	var address string
	fmt.Fscan(reader, &address)

	person := map[string]string{
		"name":    name,
		"address": address,
	}
	jsonBytes, err := json.Marshal(person)

	if err == nil {
		jsonString := string(jsonBytes)
		fmt.Println(jsonString)
	} else {
		fmt.Printf("JSON marshall failed %s\n", err)
	}
}
