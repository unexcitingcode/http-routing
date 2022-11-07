package http_routing

type Compiler[Endpoint any, Branch any, Out any] interface {
	Root(missing Endpoint) func(branches ...Branch) Out
	Path(prefix string) func(branches ...Branch) Branch
	Param(name string) func(branches ...Branch) Branch
	Get(endpoint Endpoint) Branch
	Post(endpoint Endpoint) Branch
	Put(endpoint Endpoint) Branch
	Delete(endpoint Endpoint) Branch
	Options(endpoint Endpoint) Branch
	Patch(endpoint Endpoint) Branch
	Head(endpoint Endpoint) Branch
	Connect(endpoint Endpoint) Branch
	Trace(endpoint Endpoint) Branch
}
