package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Println("Enter the name")
	fmt.Scanln(&name)
	fmt.Println("Enter the addrress")
	fmt.Scanln(&address)
	per := map[string]string{"name": name, "address": address}

	barr, err := json.Marshal(per)
	if err != nil {
		fmt.Println("error in marshal ", err)
	} else {
		fmt.Println(string(barr))
	}

}
