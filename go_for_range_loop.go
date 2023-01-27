package main

import "fmt"

func main() {

	l := []int{1, 3, 4, 5, 6, 7}

	for i, v := range l {
		fmt.Println(i, v)
	}

}
