package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
In Golang, errors are values, and there's a built-in error type called error.
Unlike many other languages, Golang doesn't have exceptions.
Instead, errors are returned as values from functions, making error handling predictable and easier to reason about.

ref. https://www.linkedin.com/pulse/mastering-error-handling-golang-shiva-raman-pandey/
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	input := scanner.Text()

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Bad String")
		return
	}

	fmt.Println(num)
}
