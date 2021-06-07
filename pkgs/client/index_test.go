package client_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PlutoClientUnitTestSuite struct {
	suite.Suite
}

func TestPlutoClientUnitTestSuite(t *testing.T) {
	suite.Run(t, new(PlutoClientUnitTestSuite))
}

func (s *PlutoClientUnitTestSuite) Test_X() {

}
