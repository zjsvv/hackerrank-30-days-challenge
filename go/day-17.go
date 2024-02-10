package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c *Calculator) power(n, p int) (int, error) {
	if n < 0 || p < 0 {
		return 0, fmt.Errorf("n and p should be non-negative")
	}

	return int(math.Pow(float64(n), float64(p))), nil

}

/*
Sample Input
4
3 5
2 4
-1 -2
-1 3

Sample Output
243
16
n and p should be non-negative
n and p should be non-negative
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	t, _ := strconv.ParseInt(scanner.Text(), 0, 64)

	// read t times from stdin
	for t > 0 {
		scanner.Scan()
		line := scanner.Text()

		arr := strings.Split(line, " ")

		n := arr[0]
		p := arr[1]

		n_int, _ := strconv.Atoi(n)
		p_int, _ := strconv.Atoi(p)

		myCalculator := &Calculator{}

		res, err := myCalculator.power(n_int, p_int)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}

		t--
	}
}
