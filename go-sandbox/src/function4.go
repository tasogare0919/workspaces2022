package main

import "fmt"

func main () {

	func(msg string){
		fmt.Println(msg)
	}("tada")

}