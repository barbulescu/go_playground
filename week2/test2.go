package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	instr := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the string : ")
	nstr, _ := instr.ReadString('\n')
	str := strings.TrimSpace(strings.ToLower(nstr))

	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
