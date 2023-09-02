package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	jsonMap := make(map[string]string)

	fmt.Print("Enter a name: ")
	var name string
	fmt.Scan(&name)

	jsonMap["name"] = name

	fmt.Print("Enter an address: ")
	var address string
	fmt.Scan(&address)

	jsonMap["address"] = address

	json, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(json))
}