package http_routing

import (
	"strings"
)

type FlatRouteTranspiler[Endpoint any, Branch any, Root any] struct {
	compiler Compiler[Endpoint, Branch, Root]
}

func NewFlatRouteTranspiler[Endpoint any, Branch any, Root any](
	compiler Compiler[Endpoint, Branch, Root],
) FlatRouteTranspiler[Endpoint, Branch, Root] {
	return FlatRouteTranspiler[Endpoint, Branch, Root]{compiler}
}

type httpMethod int64

const (
	Get httpMethod = iota
	Post
	Put
	Delete
	Options
	Patch
	Head
	Connect
	Trace
)

type FlatRoute[Endpoint any] struct {
	method   httpMethod
	path     []string
	endpoint Endpoint
}

func newFlatRoute[Endpoint any](method httpMethod, path string, endpoint Endpoint) FlatRoute[Endpoint] {
	return FlatRoute[Endpoint]{method, strings.Split(path[1:], "/"), endpoint}
}

func (flatRoute FlatRoute[Endpoint]) shift() (string, FlatRoute[Endpoint]) {
	if len(flatRoute.path) == 0 {
		return "", flatRoute
	}
	x, xs := flatRoute.path[0], flatRoute.path[1:]
	return x, FlatRoute[Endpoint]{flatRoute.method, xs, flatRoute.endpoint}

}

func groupByAndShift[Endpoint any](routes []FlatRoute[Endpoint]) map[string][]FlatRoute[Endpoint] {
	groups := make(map[string][]FlatRoute[Endpoint])
	for _, route := range routes {
		head, remaining := route.shift()
		groups[head] = append(groups[head], remaining)
	}
	return groups
}

func compileSegment[Endpoint any, Branch any, Root any](
	compiler Compiler[Endpoint, Branch, Root],
	segment string,
	routes []FlatRoute[Endpoint],
) Branch {
	groups := groupByAndShift(routes)
	branches := make([]Branch, 0, len(groups)+len(groups[""]))
	for segment, children := range groups {
		if segment == "" {
			for _, child := range children {
				compiled := compileHttpMethod(compiler, child.method, child.endpoint)
				branches = append(branches, compiled)
			}
		} else {
			compiled := compileSegment(compiler, segment, children)
			branches = append(branches, compiled)
		}
	}
	if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
		return compiler.Param(segment[1 : len(segment)-1])(branches...)
	} else {
		return compiler.Path("/" + segment)(branches...)
	}
}

func compileIndexes[Endpoint any, Branch any, Root any](
	compiler Compiler[Endpoint, Branch, Root],
	routes []FlatRoute[Endpoint],
) Branch {
	indexes := make([]Branch, 0, len(routes))
	for _, child := range routes {
		index := compileHttpMethod(compiler, child.method, child.endpoint)
		indexes = append(indexes, index)
	}
	return compiler.Path("/")(indexes...)
}

func compileHttpMethod[Endpoint any, Branch any, Root any](
	compiler Compiler[Endpoint, Branch, Root],
	method httpMethod,
	endpoint Endpoint,
) Branch {
	switch method {
	case Get:
		return compiler.Get(endpoint)
	case Post:
		return compiler.Post(endpoint)
	case Put:
		return compiler.Put(endpoint)
	case Delete:
		return compiler.Delete(endpoint)
	case Options:
		return compiler.Options(endpoint)
	case Patch:
		return compiler.Patch(endpoint)
	case Head:
		return compiler.Head(endpoint)
	case Connect:
		return compiler.Connect(endpoint)
	case Trace:
		return compiler.Trace(endpoint)
	}
	panic("invalid http method")
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Root(
	missing Endpoint,
) func(routes ...FlatRoute[Endpoint]) Root {
	return func(routes ...FlatRoute[Endpoint]) Root {
		groups := groupByAndShift(routes)
		branches := make([]Branch, 0, len(groups))
		for segment, children := range groups {
			if segment == "" {
				compiled := compileIndexes(transpiler.compiler, children)
				branches = append(branches, compiled)
			} else {
				compiled := compileSegment(transpiler.compiler, segment, children)
				branches = append(branches, compiled)
			}
		}
		return transpiler.compiler.Root(missing)(branches...)
	}
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Get(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Get, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Post(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Post, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Put(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Put, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Delete(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Delete, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Options(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Options, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Patch(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Patch, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Head(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Head, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Connect(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Connect, path, endpoint)
}

func (transpiler *FlatRouteTranspiler[Endpoint, Branch, Root]) Trace(
	path string,
	endpoint Endpoint,
) FlatRoute[Endpoint] {
	return newFlatRoute(Trace, path, endpoint)
}
