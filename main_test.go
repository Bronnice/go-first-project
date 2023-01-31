package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SuiteForCheckStringForParity struct {
	suite.Suite
}

func (suite *SuiteForCheckStringForParity) TestEmptyString() {

	strnull := ""

	result, err := checkStringForParity(strnull)

	assert.NotNil(suite.T(), result)
	assert.NotEqual(suite.T(), nil, err, "error must be not nil")
}

func (suite *SuiteForCheckStringForParity) TestNonParyString() {
	str1 := "qwe"

	result, err := checkStringForParity(str1)

	assert.Equal(suite.T(), false, *result, "result must be false")
	assert.NotNil(suite.T(), err)
}

func (suite *SuiteForCheckStringForParity) TestParyString() {
	str2 := "qwer"

	result, err := checkStringForParity(str2)

	assert.Equal(suite.T(), true, *result, "result must be true")
	assert.NotNil(suite.T(), err)
}

func SuiteForCheckStringForParityRun(t *testing.T) {
	suite.Run(t, new(SuiteForCheckStringForParity))
}
