package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(returnedDate, dueDate []string) int {
	d1, _ := strconv.Atoi(returnedDate[0])
	m1, _ := strconv.Atoi(returnedDate[1])
	y1, _ := strconv.Atoi(returnedDate[2])

	d2, _ := strconv.Atoi(dueDate[0])
	m2, _ := strconv.Atoi(dueDate[1])
	y2, _ := strconv.Atoi(dueDate[2])

	if y1 < y2 {
		return 0
	} else if y1 > y2 {
		return 10000
	}

	// y1==y2 here
	if m1 < m2 {
		return 0
	} else if m1 > m2 {
		return 500 * (m1 - m2)
	}

	// y1==y2 and m1==m2 here
	if d1 < d2 {
		return 0
	} else if d1 > d2 {
		return 15 * (d1 - d2)
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	returnedDate := strings.Split(scanner.Text(), " ")

	scanner.Scan()
	dueDate := strings.Split(scanner.Text(), " ")

	fmt.Println(calculate(returnedDate, dueDate))
}
