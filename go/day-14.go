package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Difference struct {
	elements          []int32
	MaximunDifference int32
}

func (d *Difference) computeDifference() {
	for i := 0; i < len(d.elements)-1; i++ {
		for j := i + 1; j < len(d.elements); j++ {
			currDiff := Abs(d.elements[i] - d.elements[j])
			if currDiff > d.MaximunDifference {
				d.MaximunDifference = currDiff
			}

		}
	}
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	/*
				Sample Input
				STDIN   Function
				-----   --------
				3       __elements[] size N = 3
				1 2 5   __elements = [1, 2, 5]

				Sample Output
				4

		3
		1 2 5
	*/

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)

	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32
	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	d := &Difference{elements: arr, MaximunDifference: 0}
	d.computeDifference()

	fmt.Println(d.MaximunDifference)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
