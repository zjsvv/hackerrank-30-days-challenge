package main

import (
	"bufio"
	"fmt"
	"os"
)

type Solution struct {
	stack []rune
	queue []rune
}

func (s *Solution) pushCharacter(c rune) {
	s.stack = append(s.stack, c)
}

func (s *Solution) popCharacter() (res rune) {
	res = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return
}

func (s *Solution) enqueueCharacter(c rune) {
	s.queue = append(s.queue, c)
}

func (s *Solution) dequeueCharacter() (res rune) {
	res = s.queue[0]
	s.queue = s.queue[1:len(s.queue)]
	return
}

/*
Sample Input
racecar

Sample Output
The word, racecar, is a palindrome.
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	s := []rune(input)

	p := &Solution{}

	for _, c := range s {
		p.pushCharacter(c)
		p.enqueueCharacter(c)
	}

	isPalindrome := true
	for i := 0; i < len(s)/2; i++ {
		if p.popCharacter() != p.dequeueCharacter() {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Printf("The word, %s, is a palindrome.\n", input)
	} else {
		fmt.Printf("The word, %s, is not a palindrome.\n", input)
	}

}
