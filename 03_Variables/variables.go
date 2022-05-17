package main

import "fmt"

func main() {
	var x string
	x = "first "
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	var y string = "hello"
	var z string = "world"
	var y2 string = "hello"
	fmt.Println(y == z)
	fmt.Println(y == y2)

	assignTest := "This is a test for type inferring"
	fmt.Println(assignTest)
}
