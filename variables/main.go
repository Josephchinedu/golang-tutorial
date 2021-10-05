package main

import "fmt"

func main() {
	var age int = 12
	fmt.Println("my age is: ", age)

	//** DECLARING VARIABLES **///

	// Syntax: var var_name type

	var s1 string
	s1 = "Learning Golang!"
	fmt.Println(s1)

	//** TYPE INFERENCE **//

    // Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

	var k int = 5
	var i = 10
	var j = 12.7
	fmt.Println("i:", i, "j:", j, "k:", k)
}