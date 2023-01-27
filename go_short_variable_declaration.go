package main

import "fmt"

//package level scope and var operator will set the variable
//to zero value

var y int

func main() {

	//short variable declaration operator ":="
	x := 7

	//%T will show me the type
	fmt.Println("%T", x)

	fmt.Println(y)
}
