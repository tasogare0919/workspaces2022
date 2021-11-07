package main

import "fmt"

func intSec() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSec()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSec()
	fmt.Println(newInts())
}
