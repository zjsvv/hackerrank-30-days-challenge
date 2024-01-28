package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	fmt.Println("Hello, World.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(scanner.Text())

}
