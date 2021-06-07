package base

import "sync"

type Provider interface {
	Setup()
	Run()
	Shutdown()
}

func RunProvider(wg *sync.WaitGroup, p Provider) {
	defer wg.Done()

	p.Run()
}