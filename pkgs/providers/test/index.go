package test

import (
	"log"
	"github.com/stretchr/testify/mock"
)

type Provider struct {
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