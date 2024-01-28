package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var a uint64
	var b float64
	var c string

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	a, _ = strconv.ParseUint(scanner.Text(), 0, 64)

	scanner.Scan()
	b, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	c = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + a)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+b)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + c)

}
