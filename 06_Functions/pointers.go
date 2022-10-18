package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPointer(xPtr *int) {
	*xPtr = 0

}

func main() {
	x := 5
	zero(x)
	fmt.Println("x is still: ")
	fmt.Println(x)

	zeroPointer(&x)
	fmt.Println("x is now: ")
	fmt.Println(x)
}
