package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	validNames := []string{}

	for NItr := 0; NItr < int(N); NItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstName := firstMultipleInput[0]

		emailID := firstMultipleInput[1]

		match, _ := regexp.MatchString("^([a-zA-Z0-9._%-]+@gmail\\.com)$", emailID)
		if match {
			validNames = append(validNames, firstName)
		}
	}

	sort.Slice(validNames, func(i, j int) bool {
		return validNames[i] < validNames[j]
	})

	for _, name := range validNames {
		fmt.Println(name)
	}
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
