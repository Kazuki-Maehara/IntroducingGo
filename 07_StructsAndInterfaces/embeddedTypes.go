package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

type Android_2 struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Person.Name = "Bob"
	a.Person.Talk()
	// a.Talk()

	b := new(Android_2)
	b.Name = "Tom"
	b.Talk()
	b.Person.Talk()
}
