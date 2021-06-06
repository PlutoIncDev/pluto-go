package rpc

type rpcHandlerFunc = func()

type rpcHandler struct {
	Name    string
	Handler rpcHandlerFunc
}

func newRpcHandler(name string, handler rpcHandlerFunc) *rpcHandler {
	return &rpcHandler{
		Name: name,
		Handler: handler,
	}
}

type Rpc struct {
	handlers []*rpcHandler
}

func newRpc() *Rpc {
	return new(Rpc)
}

func (r *Rpc) registerRpcHandler(name string, handler rpcHandlerFunc) {
	r.handlers = append(r.handlers, newRpcHandler(name, handler))
}

func (r *Rpc) GetHandlers() []*rpcHandler {
	return r.handlers
}

func (r *Rpc) Register(name string, handler rpcHandlerFunc) {
	r.registerRpcHandler(name, handler)
}