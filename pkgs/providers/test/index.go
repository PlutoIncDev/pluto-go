package test

import (
	"github.com/stretchr/testify/mock"
	"log"
	"pluto/pkgs/providers/base"
)

type Provider struct {
	base.Provider
	mock.Mock
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Setup() {
	log.Println("Test Provider is Setting Up")
	p.Called()
}

func (p *Provider) Run() {
	log.Println("Test Provider is Running")
	p.Called()
}

func (p *Provider) Shutdown() {
	log.Println("Test Provider is Shutting Down")
	p.Called()
}