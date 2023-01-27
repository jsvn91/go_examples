package main

import "fmt"

type person_age int
type person_name string

// the field name starts with lowercase will be not visible to outside package
// and field that starts with Cap will be visible to outside package
type persons1 struct {
	Pid    int
	pfname person_name
	plname person_name
	page   person_age
}

func main() {
	var p1 person_age
	p1 = 5

	println(p1)

	//values will be inserted according to order of fieldname
	s1 := persons1{
		1,
		"Swapnil",
		"Jagtap",
		3,
	}
	fmt.Println(s1.Pid)
}
