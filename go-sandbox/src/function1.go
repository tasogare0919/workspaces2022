package main 

import "fmt"

func hi(name string) (msg string) {
	// fmt.Println("Hello!" + name)
	msg = "Hello!" + name
	return
}

func main() {
	// hi("tada")
	fmt.Println(hi("tada"))
}