package service

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

//Suite for method IsString Even
type IsStringEvenTestSuite struct {
	suite.Suite
}

func Test_IsStringEvenTestSuite(t *testing.T) {
	suite.Run(t, new(IsStringEvenTestSuite))
}

func (suite *IsStringEvenTestSuite) Test_IsStringEven_withEmptyString_expectEror() {
	emptyString := ""

	isEven, err := IsStringEven(emptyString)

	suite.Nil(isEven)
	suite.NotNil(err)
}

func (suite *IsStringEvenTestSuite) Test_IsStringEven_OddString_expectOdd() {
	oddString := "qwe"

	isEven, err := IsStringEven(oddString)

	suite.NotNil(*isEven)
	suite.Nil(err)
}

func (suite *IsStringEvenTestSuite) Test_IsStringEven_EvenString_expectEven() {
	evenString := "qwer"

	isEven, err := IsStringEven(evenString)

	suite.NotNil(*isEven)
	suite.Nil(err)
}
