package service

import (
	"errors"

	"stringvalidator/utils"
)

func IsStringEven(str string) (*bool, error) {
	if len(str) == 0 {
		return nil, errors.New("lenght of the string is 0")
	}
	return utils.GetPointer(len(str)%2 == 0), nil
}
