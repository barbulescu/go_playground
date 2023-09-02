package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter a float number:")
	reader := bufio.NewReader(os.Stdin)
	var floatValue float32
	fmt.Fscan(reader, &floatValue)
	var intValue = int(floatValue)
	fmt.Printf("%f was converted to %d\n", floatValue, intValue)
}
