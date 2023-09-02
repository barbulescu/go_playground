package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fName string
	lName string
}

func main() {
	names := make([]Name, 0, 10)
	var fileName string
	fmt.Println("Enter the file name to read:")
	fmt.Scanln(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error in reading file", err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			names = append(names, Name{fName: words[0], lName: words[1]})
		}

		for _, name := range names {
			fmt.Printf("First name %s, Last Name %s", name.fName, name.lName)
			fmt.Println()
		}
	}
}
