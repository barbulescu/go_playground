package main

import "fmt"

func main() {
	var float float64

	fmt.Print("Write a float number: ")
	fmt.Scan(&float)

	var parsed int = int(float)

	fmt.Printf("Integer: %d\n", parsed)
}
