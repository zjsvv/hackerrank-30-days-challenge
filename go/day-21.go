package main

import "fmt"

type T interface {
	int | string
}

func printArray[t T](arr []t) {
	for _, element := range arr {
		fmt.Println(element)
	}
}

func main() {
	var n int
	var intArr []int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		intArr = append(intArr, number)
	}

	var strArr []string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		strArr = append(strArr, s)
	}

	printArray(intArr)
	printArray(strArr)
}
