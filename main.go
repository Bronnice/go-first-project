package main

import (
	"errors"
)

func checkStringForParity(str string) (*bool, error) {
	if len(str) == 0 {
		return nil, errors.New("lenght of the string is 0")
	}
	if len(str)%2 != 0 {
		tempbool := false
		return &tempbool, nil
	} else {
		tempbool := true
		return &tempbool, nil
	}
}

func main() {
}
