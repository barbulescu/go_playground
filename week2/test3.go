package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var found bool

func main() {
	var strInput string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strInput = scanner.Text()
	}
	nameLower := strings.ToLower(strInput)
	nameSpace := strings.ReplaceAll(nameLower, " ", "")

	if strings.HasPrefix(nameSpace, "i") && strings.HasSuffix(nameSpace, "n") && strings.Contains(nameSpace, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
