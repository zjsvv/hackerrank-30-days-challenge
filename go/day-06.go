package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	numberOfStrings, _ := strconv.ParseInt(scanner.Text(), 0, 64)

	var i int64

	for i < numberOfStrings {
		scanner.Scan()
		currStr := scanner.Text()

		evenChars := ""
		oddChars := ""

		for j := 0; j < len(currStr); j++ {
			if j%2 == 0 {
				evenChars += string(currStr[j])
			} else {
				oddChars += string(currStr[j])
			}
		}

		fmt.Printf("%s %s\n", evenChars, oddChars)

		i++
	}
}
