package main

import (
	"flag"
	"fmt"

	"qwer/service"
)

func main() {
	flagString := flag.String("mystring", "", "Your own string")
	flag.Parse()

	result, err := service.IsStringEven(*flagString)
	if err != nil {
		fmt.Println(err)
	} else {
		if *result {
			fmt.Println("Even number of letters")
		} else {
			fmt.Println("Odd number of letters")
		}
	}
}
