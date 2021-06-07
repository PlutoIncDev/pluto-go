package providers

import (
	"pluto/pkgs/providers/base"
	"sync"
)

func RunProvider(wg *sync.WaitGroup, p base.Provider) {
	defer wg.Done()

	p.Run()
}