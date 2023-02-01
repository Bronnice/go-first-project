package service

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IsStringEvenTestSuite struct {
	suite.Suite
}

func Test_IsStringEvenTestSuite(t *testing.T) {
	suite.Run(t, new(IsStringEvenTestSuite))
}

//Test of method IsStringEven with empty string
func (suite *IsStringEvenTestSuite) Test_IsStringEven_withEmptyString() {
	emptyString := ""

	result, err := IsStringEven(emptyString)

	suite.NotEqual(result, nil)
	suite.NotEqual(err, nil)
}

//Test of method IsStringEven with odd string
func (suite *IsStringEvenTestSuite) Test_IsStringEven_OddString() {
	oddString := "qwe"

	result, err := IsStringEven(oddString)

	suite.Equal(*result, false)
	suite.Equal(err, nil)
}

//Test of method IsStringEven with even string
func (suite *IsStringEvenTestSuite) Test_IsStringEven_EvenString() {
	evenString := "qwer"

	result, err := IsStringEven(evenString)

	suite.Equal(*result, true)
	suite.Equal(err, nil)
}
