package main

import "fmt"

func SimpleFunction() {

	fmt.Println("Hi I am a functions")
}

func SimpleAdd(x int, y int) (total int) {

	total = x + y
	// Return statement without specify variable name
	return
}

type persons struct {
	pid    int
	pfname string
	plname string
}

// attach a reciver to a function
func (p persons) Speak() {
	fmt.Println("Hello ", p.pfname)
}

func main() {
	SimpleFunction()
	total := SimpleAdd(2, 3)
	fmt.Println("the addition of 2 + 3 is ", total)

	var p persons
	p = persons{
		1,
		"Swapnil",
		"Jagtap",
	}
	p.Speak()
}
