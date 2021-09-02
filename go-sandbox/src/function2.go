package main

import "fmt"

func swap(a, b int) (int, int) {
	return b,a
}

func main() {
	fmt.Println(swap(6, 7))
}