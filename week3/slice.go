package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	x := [...]int{1, 2, 3, 4}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))

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
