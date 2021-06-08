package rpc

import (
	"context"
	"log"
	"time"
)

func (p *Provider) Setup() {
	log.Println("Setting up RPC Server")
}

func (p *Provider) Run(ctx context.Context) {
	log.Println("Running RPC Server")
}

func (p *Provider) Shutdown() {
	time.Sleep(time.Second * 3) // simulated delay.. TODO: remove
	p.finished = true
	log.Println("Finished shutting down RPC Server")
}

func (p *Provider) IsFinished() bool {
	return p.finished
}
