package main

import (
	"fmt"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var test string
	test = "Hello World"
	w.Write([]byte(test))
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	fmt.Println("Web Server Start!")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error")
	}
}
