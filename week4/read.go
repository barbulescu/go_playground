package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Name struct {
	fname string
	lname string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter filename: ")
	var filename string
	fmt.Fscan(reader, &filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var results = []Name{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fname := line[0:20]
		lname := line[21:41]
		name := Name{fname, lname}
		results = append(results, name)
	}

	for _, name := range results {
		fmt.Printf("fname: %s, lname: %s\n", name.fname, name.lname)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
