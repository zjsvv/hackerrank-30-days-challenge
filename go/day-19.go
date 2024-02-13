package main

import (
	"fmt"
	"math"
	"reflect"
)

type AdvancedArithmetic interface {
	divisorSum(n int) int
}

type Calculator struct{}

func (c *Calculator) divisorSum(n int) int {
	if n == 1 {
		return 1
	}

	res := 0

	sqrtRoot := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrtRoot; i++ {
		if n%i == 0 {
			res += i + n/i
		}
	}

	if sqrtRoot*sqrtRoot == n {
		res -= sqrtRoot
	}

	return res
}

/*
Sample Input
6

Sample Output
I implemented: AdvancedArithmetic
12

*/

func main() {
	var n int

	fmt.Scan(&n)

	var myCalculator AdvancedArithmetic = &Calculator{}

	sum := myCalculator.divisorSum(n)

	c := reflect.TypeOf(myCalculator)
	a := reflect.TypeOf((*AdvancedArithmetic)(nil)).Elem()
	if c.Implements(a) {
		fmt.Println("I implemented:  " + a.Name())
	}

	// if _, ok := myCalculator.(AdvancedArithmetic); ok {
	// 	fmt.Println("I implemented:  " + reflect.TypeOf((*AdvancedArithmetic)(nil)).Elem().Name())
	// }

	fmt.Println(sum)
}
