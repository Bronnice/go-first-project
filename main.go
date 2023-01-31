package main

import (
	"errors"
	"flag"
	"fmt"
)

func getPointer[T any](obj T) *T {
	return &obj
}

func IsStringEven(str string) (*bool, error) {
	if len(str) == 0 {
		return nil, errors.New("lenght of the string is 0")
	}
	if len(str)%2 != 0 {
		return getPointer(false), nil
	}
	return getPointer(true), nil
}

func main() {
	flagString := flag.String("mystring", "", "Your own string")
	flag.Parse()

	result, err := IsStringEven(*flagString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*result)
	}
}
