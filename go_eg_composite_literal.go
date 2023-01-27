package main

import "fmt"

func main() {
	//this is slice which is used to hold a list
	//and type{} is called composite literal , basically type{}
	xi := []int{1, 2, 3, 4}

	fmt.Println(xi)

	m := map[string]int{
		"one": 2,
		"two": 3,
	}

	fmt.Println(m["one"])
}
