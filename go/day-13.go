package main

import "fmt"

type Book struct {
	title  string
	author string
}

type MyBook struct {
	Book
	price int
}

func (mb *MyBook) display() {
	fmt.Printf("Title: %s\n", mb.title)
	fmt.Printf("Author: %s\n", mb.author)
	fmt.Printf("Price: %d\n", mb.price)
}

func main() {
	var title string
	var author string
	var price int

	fmt.Scan(&title)
	fmt.Scan(&author)
	fmt.Scan(&price)

	mb := &MyBook{Book{title, author}, price}
	mb.display()
}
