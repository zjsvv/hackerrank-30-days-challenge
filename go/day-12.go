package main

import (
	"fmt"
)

type Person struct {
	firstName      string
	lastName       string
	identification int
}

func (p *Person) printPerson() {
	fmt.Printf("Name: %s, %s\nID: %d\n", p.lastName, p.firstName, p.identification)
}

type Student struct {
	Person
	scores []int
}

func (s *Student) calculate() string {
	total := 0
	for _, score := range s.scores {
		total += score
	}

	avgScore := total / len(s.scores)

	switch {
	case avgScore >= 90 && avgScore <= 100:
		return "O"
	case avgScore >= 80 && avgScore < 90:
		return "E"
	case avgScore >= 70 && avgScore < 80:
		return "A"
	case avgScore >= 50 && avgScore < 70:
		return "P"
	case avgScore >= 40 && avgScore < 55:
		return "D"
	default:
		return "T"
	}
}

/*
Input (stdin)
Heraldo Memelli 8135627
2
100 80


Expected Output
Name: Memelli, Heraldo
ID: 8135627
Grade: O
*/

func main() {
	var firstName string
	var lastName string
	var id int
	var scoresLength int

	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&id)
	fmt.Scan(&scoresLength)

	scores := []int{}
	for i := 0; i < scoresLength; i++ {
		var score int
		fmt.Scan(&score)
		scores = append(scores, score)
	}

	student := &Student{Person: Person{firstName, lastName, id}, scores: scores}

	student.printPerson()
	fmt.Printf("Grade: %s\n", student.calculate())
}
