package internal_test

import (
	"github.com/stretchr/testify/suite"
	"pluto/pkgs/internal"
	"testing"
)


type InternalOptionsUnitTestSuite struct {
	suite.Suite
}

func TestInternalOptionsUnitTestSuite(t *testing.T) {
	suite.Run(t, new(InternalOptionsUnitTestSuite))
}

func (s *InternalOptionsUnitTestSuite) Test_CreateOptions_Success() {
	opts, err := internal.NewOptions("test")

	s.NoError(err)
	s.NotNil(opts)
}

func (s *InternalOptionsUnitTestSuite) Test_CreateOptions_NameTest() {
	tests := map[string]struct {
		input string
		shouldErr bool
	}{
		"len 0": {
			input: "",
			shouldErr: true,
		},
		"len 1": {
			input: "a",
			shouldErr: true,
		},
		"len 2": {
			input: "ab",
			shouldErr: true,
		},
		"len 3": {
			input: "abc",
			shouldErr: false,
		},
		"len 4": {
			input: "abcd",
			shouldErr: false,
		},
		"default": {
			input: "default",
			shouldErr: false,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			_, err := internal.NewOptions( tc.input)

			if tc.shouldErr && err == nil {
				s.Fail("Should error with input", tc.input)
			}

			if !tc.shouldErr && err != nil {
				s.Fail("Should not error with input", tc.input)
			}
		})
	}
}


func (s *InternalOptionsUnitTestSuite) Test_Options_GetName() {
	opts, _ := internal.NewOptions("test")

	s.Equal("test", opts.GetName())
}