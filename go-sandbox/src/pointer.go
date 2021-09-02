package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a // &a = aのアドレス
	fmt.Println(*pa)
}