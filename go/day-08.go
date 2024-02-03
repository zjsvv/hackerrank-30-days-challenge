package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.ParseInt(scanner.Text(), 0, 64)

	// read n times from stdin and update hash table
	for i := int64(0); i < n; i++ {
		scanner.Scan()
		line := scanner.Text()

		res := strings.Split(line, " ")
		m[res[0]] = res[1]
	}

	// read the rest lines from stdin
	for scanner.Scan() {
		name := scanner.Text()

		if number, exist := m[name]; exist {
			fmt.Printf("%s=%s\n", name, number)
		} else {
			fmt.Println("Not found")
		}
	}
}
