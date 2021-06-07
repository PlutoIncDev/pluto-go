package http_test

import (
	"github.com/stretchr/testify/suite"
	"pluto/pkgs/providers/http"
	"testing"
)

type ProvidersHttpIndexUnitTestSuite struct {
	suite.Suite
}

func TestProvidersHttpIndexUnitTestSuite(t *testing.T) {
	suite.Run(t, new(ProvidersHttpIndexUnitTestSuite))
}

func (s *ProvidersHttpIndexUnitTestSuite) Test_RegisterEndpoint_Success() {
	// Tests that RegisterEndpoint function is successful
	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterEndpoint(http.GetMethod, "/health-check", func() {})

	e := httpProvider.GetEndpoints()

	s.Equal(1, len(e))
	s.Equal(http.GetMethod, e[0].GetMethod())
	s.Equal("/health-check", e[0].GetPath())
}

func (s *ProvidersHttpIndexUnitTestSuite) Test_RegisterEndpoint_Duplicate_Failure() {
	// Tests that RegisterEndpoint function fails if an existing endpoint already exists
	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterEndpoint(http.GetMethod, "/health-check", func() {})

	s.Panics(func() {
		httpProvider.RegisterEndpoint(http.GetMethod, "/health-check", func() {})
	})
}

func (s *ProvidersHttpIndexUnitTestSuite) Test_GetEndpoints() {
	// Tests that the GetEndpoints function returns the correct endpoints
	httpProvider := http.NewProvider("8080")

	httpProvider.RegisterEndpoint(http.GetMethod, "/health-check", func() {})
	httpProvider.RegisterEndpoint(http.PostMethod, "/user/auth", func() {})

	e := httpProvider.GetEndpoints()
	s.Equal(2, len(e))

	s.Equal(http.GetMethod, e[0].GetMethod())
	s.Equal("/health-check", e[0].GetPath())
	s.Equal(http.PostMethod, e[1].GetMethod())
	s.Equal("/user/auth", e[1].GetPath())

}
