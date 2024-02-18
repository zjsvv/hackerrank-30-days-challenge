package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	sqrtRoot := int(math.Sqrt(float64(num)))

	for i := 2; i < sqrtRoot+1; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	var n int

	fmt.Scan(&n)

	for n > 0 {
		var num int
		fmt.Scan(&num)

		if isPrime(num) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}

		n--
	}
}
