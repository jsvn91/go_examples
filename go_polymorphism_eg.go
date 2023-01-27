package main

import "fmt"

type person struct {
	pid    int
	pfname string
	plname string
}

type secret_agent struct {
	person
}

func (p person) speak() {
	fmt.Println("Hi, I am ", p.pfname)
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		pid:    1,
		pfname: "Swapnil",
		plname: "Jagtap",
	}

	s1 := secret_agent{
		person{
			pid:    2,
			pfname: "Rohit",
			plname: "Jagtap",
		},
	}

	saysomething(p1)
	saysomething(s1)
}
