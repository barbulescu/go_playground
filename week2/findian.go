package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter a string: ")
	in := bufio.NewReader(os.Stdin)
	value, err := in.ReadString('\n')

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	value = strings.ToLower(value)
	value = strings.TrimSuffix(value, "\n")

	if strings.HasPrefix(value, "i") && strings.HasSuffix(value, "n") && strings.Contains(value, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
