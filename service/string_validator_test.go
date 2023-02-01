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

func (suite *IsStringEvenTestSuite) Test_IsStringEven_withEmptyString() {
	emptyString := ""

	result, err := IsStringEven(emptyString)

	suite.Nil(result)
	suite.NotNil(err)
}

func (suite *IsStringEvenTestSuite) Test_IsStringEven_OddString() {
	oddString := "qwe"

	result, err := IsStringEven(oddString)

	suite.Equal(*result, false)
	suite.Nil(err)
}

func (suite *IsStringEvenTestSuite) Test_IsStringEven_EvenString() {
	evenString := "qwer"

	isEven, err := IsStringEven(evenString)

	suite.Equal(*isEven, true)
	suite.Nil(err)
}
