package pluto_test

import (
	"github.com/stretchr/testify/suite"
	"pluto/pkgs/pluto"
	"pluto/pkgs/testutils"
	"testing"
)

type ClientUnitTestSuite struct {
	suite.Suite
}

func TestClientUnitTestSuite(t *testing.T) {
	suite.Run(t, new(ClientUnitTestSuite))
}

func (s *ClientUnitTestSuite) Test_CreateNewClient_Success() {
	c, err := pluto.NewClient("test")
	s.NoError(err)
	s.NotNil(c)
	s.Equal(c.Options.GetName(), "test")
}

func (s *ClientUnitTestSuite) Test_CreateNewClient_NameError() {
	c, err := pluto.NewClient("")

	s.Error(err)
	s.Nil(c)
}

func (s *ClientUnitTestSuite) Test_ClientString() {
	c, _ := pluto.NewClient("test")

	// Assert that the string generated when printing the Client is correct
	s.Equal("<PlutoClient name='test'/>", c.String())
}

func (s *ClientUnitTestSuite) Test_ClientRegisterEntrypoint() {
	c, _ := pluto.NewClient("test")

	testEntrypoint := testutils.NewMockTestEntrypoint()

	testEntrypoint.On("Setup").Return(nil)

	// Assert we start with 0 entrypoints defined
	s.Equal(0, len(c.Entrypoints))

	// Register the Test Entrypoint
	c.RegisterEntrypoint(&testEntrypoint)

	// Assert we insert the test entrypoint
	s.Equal(1, len(c.Entrypoints))
	s.Equal(&testEntrypoint, c.Entrypoints[0])

	// assert that Setup is called when RegisterEntrypoint is called
	testEntrypoint.AssertNumberOfCalls(s.T(), "Setup", 1)
}