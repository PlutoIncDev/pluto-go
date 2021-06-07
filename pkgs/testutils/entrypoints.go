package testutils

import "github.com/stretchr/testify/mock"

type MockTestEntrypoint struct {
	mock.Mock
}

func NewMockTestEntrypoint() MockTestEntrypoint {
	return MockTestEntrypoint{}
}

func (m *MockTestEntrypoint) Setup() {
	m.Called()
}

func (m *MockTestEntrypoint) Run() {
	m.Called()
}

func (m MockTestEntrypoint) Shutdown() {
	m.Called()
}