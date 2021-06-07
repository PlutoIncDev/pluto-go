package base

type Provider struct {

}

/* Overridable Functions */
func (p *Provider) Setup() {
	panic("Not Implemented Setup for Provider")
}

func (p *Provider) Run() {
	panic("Not Implemented Run Setup for Provider")
}

func (p *Provider) Shutdown() {
	panic("Not Implemented Shutdown for Provider")
}

/* Internal Functions (ideally not messed with) */
func (p *Provider) ISetup() {
	p.Setup()
}

func (p *Provider) IRun() {
	p.Run()
}

func (p *Provider) IShutdown() {
	p.Shutdown()
}
