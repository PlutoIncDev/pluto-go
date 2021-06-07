package pluto

import (
	"fmt"
	"log"
	"pluto/pkgs/providers"
	"sync"
)

type Client struct {
	name string
	providers []providers.Provider
}

func NewClient(name string) *Client {
	if len(name) == 0 {
		panic("Name of service must be at least 1 character long")
	}

	return &Client{name: name}
}

func (c *Client) RegisterProvider(p providers.Provider) {
	c.providers = append(c.providers, p)

	p.ISetup()
}

func (c *Client) GetProviders() []providers.Provider {
	return c.providers
}

func (c *Client) Start() {
	log.Println(fmt.Sprintf("Starting the %s service", c.name))

	var wg sync.WaitGroup

	for _, p := range c.providers {
		wg.Add(1)
		go providers.RunProvider(&wg, p)
	}

	wg.Wait()

	for _, p := range c.providers {
		p.IShutdown()
	}
}
