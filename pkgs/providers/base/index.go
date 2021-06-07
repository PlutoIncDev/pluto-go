package base

type Provider interface {
	Setup()
	Run()
	Shutdown()
}
