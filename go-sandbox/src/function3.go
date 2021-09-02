package main

import "fmt"

func main () {
	f := func(a, b int) (int, int) {
		return a, b
	}

	fmt.Println(f(3,2))
}