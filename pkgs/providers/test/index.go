package test

import (
	"github.com/stretchr/testify/mock"
)

type Provider struct {
	mock.Mock
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Setup() {
	p.Called()
}

func (p *Provider) Run() {
	p.Called()
}

func (p *Provider) Shutdown() {
	p.Called()
}