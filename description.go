package http_routing

type DescriptionCompiler[Endpoint any] struct{}

func NewDescriptionCompiler[Endpoint any]() Compiler[Endpoint, []RouteDescription[Endpoint], Description[Endpoint]] {
	return DescriptionCompiler[Endpoint]{}
}

type Description[Endpoint any] struct {
	Missing Endpoint
	Routes  []RouteDescription[Endpoint]
}

type RouteDescription[Endpoint any] struct {
	Method   string
	Path     string
	Endpoint Endpoint
}

func (description RouteDescription[Endpoint]) prefixedWithPath(prefix string) RouteDescription[Endpoint] {
	return RouteDescription[Endpoint]{
		Method:   description.Method,
		Path:     prefix + description.Path,
		Endpoint: description.Endpoint,
	}
}

func (description RouteDescription[Endpoint]) prefixedWithParam(name string) RouteDescription[Endpoint] {
	return RouteDescription[Endpoint]{
		Method:   description.Method,
		Path:     "/{" + name + description.Path + "}",
		Endpoint: description.Endpoint,
	}
}

func prefixBranches[Endpoint any](
	prefix func(description RouteDescription[Endpoint]) RouteDescription[Endpoint],
) func(branches ...[]RouteDescription[Endpoint]) []RouteDescription[Endpoint] {
	return func(branches ...[]RouteDescription[Endpoint]) []RouteDescription[Endpoint] {
		return flattenThenMap(branches, prefix)
	}
}

func (describer DescriptionCompiler[Endpoint]) Root(
	missing Endpoint,
) func(branches ...[]RouteDescription[Endpoint]) Description[Endpoint] {
	return func(branches ...[]RouteDescription[Endpoint]) Description[Endpoint] {
		return Description[Endpoint]{Missing: missing, Routes: flatten(branches)}
	}
}

func (describer DescriptionCompiler[Endpoint]) Path(
	prefix string,
) func(branches ...[]RouteDescription[Endpoint]) []RouteDescription[Endpoint] {
	return prefixBranches(func(description RouteDescription[Endpoint]) RouteDescription[Endpoint] {
		return description.prefixedWithPath(prefix)
	})
}

func (describer DescriptionCompiler[Endpoint]) Param(
	name string,
) func(branches ...[]RouteDescription[Endpoint]) []RouteDescription[Endpoint] {
	return prefixBranches(func(description RouteDescription[Endpoint]) RouteDescription[Endpoint] {
		return description.prefixedWithParam(name)
	})
}

func (describer DescriptionCompiler[Endpoint]) Get(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"GET", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Post(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"POST", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Put(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"PUT", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Delete(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"DELETE", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Options(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"OPTIONS", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Patch(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"PATCH", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Head(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"HEAD", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Connect(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"CONNECT", "", endpoint}}
}

func (describer DescriptionCompiler[Endpoint]) Trace(endpoint Endpoint) []RouteDescription[Endpoint] {
	return []RouteDescription[Endpoint]{{"TRACE", "", endpoint}}
}
