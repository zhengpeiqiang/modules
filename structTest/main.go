package main

import "fmt"

type address struct {
	a int
}

type this interface {
	memory()
}

func (ad address) memory() {
	fmt.Println("a - ", ad)
	fmt.Println("a's memory address --> ", &ad)
}

func main() {
	ad := 43
	fmt.Println("a - ", ad)
	fmt.Println("a's memory address --> ", &ad)

	//code init in here
	thisAddress := address{
		a: 42,
	}
	// not sure why this doesnt return memory address as well?
	var i this
	i = thisAddress
	i.memory()
}