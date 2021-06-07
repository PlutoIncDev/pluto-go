package utils_test

import (
	"github.com/stretchr/testify/suite"
	"pluto/pkgs/internal/utils"
	"testing"
)

type InternalUtilsSliceUnitTestSuite struct {
	suite.Suite
}

func TestInternalUtilsSliceUnitTestSuite(t *testing.T) {
	suite.Run(t, new(InternalUtilsSliceUnitTestSuite))
}

func (s *InternalUtilsSliceUnitTestSuite) Test_StringInSlice() {
	tests := map[string]struct {
		inputStr string
		inputSli []string
		result bool
	}{
		"Single Case": {
			inputStr: "test",
			inputSli: []string{"test"},
			result: true,
		},
		"Empty Slice": {
			inputStr: "test",
			inputSli: []string{},
			result: false,
		},
		"Empty String": {
			inputStr: "",
			inputSli: []string{"test"},
			result: false,
		},
		"Empty String and Slice": {
			inputStr: "",
			inputSli: []string{},
			result: false,
		},
		"Empty String and Matching Slice": {
			inputStr: "",
			inputSli: []string{""},
			result: true,
		},
		"Multiple Cases": {
			inputStr: "test",
			inputSli: []string{"test", "one", "two"},
			result: true,
		},
		"Not Found Case": {
			inputStr: "test",
			inputSli: []string{"one", "two"},
			result: false,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.Equal(tc.result, utils.StringInSlice(tc.inputStr, tc.inputSli))
		})
	}
}


func (s *InternalUtilsSliceUnitTestSuite) Test_SlicesContainSameElement() {
	tests := map[string]struct {
		inputL1 []string
		inputL2 []string
		result bool
	}{
		"Empty Case": {
			inputL1: []string{},
			inputL2: []string{},
			result: false,
		},
		"True Case": {
			inputL1: []string{"test", "random", "values"},
			inputL2: []string{"abc", "test"},
			result: true,
		},
		"False Case": {
			inputL1: []string{"test", "random", "values"},
			inputL2: []string{"abc", "more"},
			result: false,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.Equal(tc.result, utils.SlicesContainSameElement(tc.inputL1, tc.inputL2))
		})
	}
}
