package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")
	str := recover() // This will never happen
	fmt.Println(str)
}
