package main

import "fmt"

func main() {
	const greet string = "Hello World!"
	var name string
	var age uint
	//uint is unsigned integer that can take only positive integers and int can take positive and negative values both

	name = "Kshitij Jain"
	age = 19
	fmt.Printf("%v, My name is %v and age is %v\n", greet, name, age)
}
