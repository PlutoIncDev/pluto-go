package internal

type Events struct {
	Http *Http
	Rpc *Rpc
}

func NewEvents() *Events {
	return &Events{
		Http: newHttp(),
		Rpc: newRpc(),
	}
}