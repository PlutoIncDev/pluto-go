package providers

import "sync"

type Provider interface {
	ISetup()
	Setup()

	IRun()
	Run()

	IShutdown()
	Shutdown()
}

func RunProvider(wg *sync.WaitGroup, p Provider) {
	defer wg.Done()

	p.IRun()
}