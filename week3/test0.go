package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := make([]int, 0, 3)

	for {
		fmt.Println("Please enter an integer: ")
		var value string
		fmt.Fscan(reader, &value)
		if value == "X" {
			break
		}
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
			continue
		}
		s = append(s, intValue)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		fmt.Printf("%d\n", s)
	}
}
