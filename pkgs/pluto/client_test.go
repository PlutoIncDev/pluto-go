package pluto_test

import (
	"github.com/stretchr/testify/suite"
	"pluto/pkgs/pluto"
	"pluto/pkgs/providers/test"
	"testing"
)

type PlutoClientUnitTestSuite struct {
	suite.Suite
}

func TestPlutoClientUnitTestSuite(t *testing.T) {
	suite.Run(t, new(PlutoClientUnitTestSuite))
}

func (s *PlutoClientUnitTestSuite) Test_NewClient_Test_Success() {
	// Tests that we can create a new pluto client with name 'test'
	c := pluto.NewClient("test")
	s.NotNil(c)
}

func (s *PlutoClientUnitTestSuite) Test_NewClient_T_Success() {
	// Tests that we can create a new pluto client with name 't'
	c := pluto.NewClient("t")
	s.NotNil(c)
}

func (s *PlutoClientUnitTestSuite) Test_NewClient_EmptyNamePanics() {
	// Tests that we panic if the name is not at least 1 character long
	s.Panics(func() {
		pluto.NewClient("")
	})
}

func (s *PlutoClientUnitTestSuite) Test_Client_RegisterProvider() {
	// Tests that we can register a provider successfully
	c := pluto.NewClient("test")

	testProvider := test.NewProvider()
	testProvider.On("Setup").Return(nil)

	c.RegisterProvider(testProvider)

	// assert that the providers was registered
	s.Equal(1, len(c.GetProviders()))

	// assert that Setup was called
	testProvider.AssertNumberOfCalls(s.T(), "Setup", 1)
}

func (s *PlutoClientUnitTestSuite) Test_Client_ProviderLifeCycle() {
	// Tests that the client can be started and the provider life cycle is correct

	c := pluto.NewClient("test")

	testProvider := test.NewProvider()
	testProvider.On("Setup").Return(nil)
	testProvider.On("Run").Return(nil)
	testProvider.On("Shutdown").Return(nil)

	c.RegisterProvider(testProvider)

	c.Start()

	testProvider.AssertNumberOfCalls(s.T(), "Setup", 1)
	testProvider.AssertNumberOfCalls(s.T(), "Run", 1)
	testProvider.AssertNumberOfCalls(s.T(), "Shutdown", 1)
}
