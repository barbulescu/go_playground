// slice.go

/**
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	var slice []int = make([]int, 3)

	count := 0
	for {
		fmt.Println("Please enter your integer: ")
		fmt.Scan(&input)
		if input == "X" {
			fmt.Println("Exiting...")
			break
		}

		num, _ := strconv.Atoi(input)
		// if less than 3, replace the 0 with the new number
		if count < 3 {
			// sort the slice reverse
			// sort ints descending
			sort.Sort(sort.Reverse(sort.IntSlice(slice)))
			slice[count] = num
		} else {
			slice = append(slice, num)
		}
		sort.Ints(slice)
		fmt.Println(slice)
		count++
	}
}
