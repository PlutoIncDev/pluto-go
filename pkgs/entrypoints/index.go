package entrypoints

import "sync"

type Entrypoint interface {
	/*
		Setup is called when the Entrypoint is registered. You should do
		any required setup for the entrypoint to work here.
	*/
	Setup()

	/*
		Shutdown is called before the server stops. You should do any cleanup
		needed for your extension here.
	*/
	Shutdown()

	/*
	TODO: WRITE RUN
	 */
	Run()
}

func RunEntrypoint(wg *sync.WaitGroup, e Entrypoint) {
	defer wg.Done()

	e.Run()
}

