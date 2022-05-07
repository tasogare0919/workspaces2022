package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	// var convertNumber int
	// var inputNumber string = flag.Args()
	// convertNumber := strconv.Atoi(inputNumber)
	random, _ := MakeRandomStr(16)
	fmt.Println(random)
}

func MakeRandomStr(digit int) (string, error) {
	// func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randStrings := make([]byte, digit)
	if _, err := rand.Read(randStrings); err != nil {
		return "", errors.New("unexpected error")
	}

	var result string
	for _, v := range randStrings {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
