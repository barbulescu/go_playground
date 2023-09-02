package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//Definition of a Name Struct
type name struct {
	fname string
	lname string
}

func main() {
	fileNames := make([]name, 0, 3)
	// Prompt user for a txt file
	fmt.Println("Enter the name of the text file")
	inputFileReader := bufio.NewReader(os.Stdin)
	fileName, err := inputFileReader.ReadString('\n')
	fileName = strings.TrimSuffix(fileName, "\n")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("'%s'", fileName)

	// Open a txt file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// Scan the content of a file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitNames := strings.Split(scanner.Text(), " ")
		var sliceStruct name
		sliceStruct.fname, sliceStruct.lname = splitNames[0], splitNames[1]
		fileNames = append(fileNames, sliceStruct)
	}
	// Iterate through slice of structs and print the first and last names found in each struct
	fmt.Println("First and last names found in each struct:")
	for _, iterateStruct := range fileNames {
		fmt.Println(iterateStruct.fname, iterateStruct.lname)
	}
}
