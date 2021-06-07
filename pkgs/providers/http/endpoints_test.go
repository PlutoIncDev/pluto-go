package http_test

import (
	"github.com/stretchr/testify/suite"
	"pluto/pkgs/providers/http"
	"testing"
)

type ProvidersHttpEndpointsUnitTestSuite struct {
	suite.Suite
}

func TestProvidersHttpEndpointsUnitTestSuite(t *testing.T) {
	suite.Run(t, new(ProvidersHttpEndpointsUnitTestSuite))
}

func (s *ProvidersHttpEndpointsUnitTestSuite) Test_GetMethod() {
	// Tests that the GetMethod function returns the correct method
	e := http.NewEndpoint(http.GetMethod, "/health-check", func() {})

	s.Equal(http.GetMethod, e.GetMethod())
}

func (s *ProvidersHttpEndpointsUnitTestSuite) Test_GetPath() {
	// Tests that the GetPath function returns the correct path
	e := http.NewEndpoint(http.GetMethod, "/health-check", func() {})

	s.Equal("/health-check", e.GetPath())
}
