package httpserver_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type InternalHTTPServerUnitTestSuite struct {
	suite.Suite
}

func TestInternalHTTPServerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(InternalHTTPServerUnitTestSuite))
}

func (s *InternalHTTPServerUnitTestSuite) Test_CreateHTTPServer_Success() {

}
