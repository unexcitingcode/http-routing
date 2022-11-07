package http_routing

import (
	"strings"
)

type RequestLineCompiler[Endpoint any] struct{}

type RequestLine struct {
	Method string
	Path   string
}

type RequestLineMatch[Endpoint any] struct {
	Endpoint Endpoint
	Params   map[string]string
}

func (match *RequestLineMatch[Endpoint]) withParam(key, value string) *RequestLineMatch[Endpoint] {
	newParams := map[string]string{}
	for k, v := range match.Params {
		newParams[k] = v
	}
	newParams[key] = value
	return &RequestLineMatch[Endpoint]{match.Endpoint, newParams}
}

type RequestLineBranch[Endpoint any] func(method string, remaining string) *RequestLineMatch[Endpoint]
type RequestLineRoot[Endpoint any] func(line RequestLine) RequestLineMatch[Endpoint]

func NewRequestLineCompiler[Endpoint any]() Compiler[Endpoint, RequestLineBranch[Endpoint], RequestLineRoot[Endpoint]] {
	return RequestLineCompiler[Endpoint]{}
}

func (compiler RequestLineCompiler[Endpoint]) Root(
	missing Endpoint,
) func(branches ...RequestLineBranch[Endpoint]) RequestLineRoot[Endpoint] {
	return func(branches ...RequestLineBranch[Endpoint]) RequestLineRoot[Endpoint] {
		return func(line RequestLine) RequestLineMatch[Endpoint] {
			for _, branch := range branches {
				match := branch(line.Method, line.Path)
				if match != nil {
					return *match
				}
			}
			return RequestLineMatch[Endpoint]{missing, map[string]string{}}
		}
	}
}

func (compiler RequestLineCompiler[Endpoint]) Path(
	prefix string,
) func(branches ...RequestLineBranch[Endpoint]) RequestLineBranch[Endpoint] {
	return func(branches ...RequestLineBranch[Endpoint]) RequestLineBranch[Endpoint] {
		return func(method string, remaining string) *RequestLineMatch[Endpoint] {
			if !strings.HasPrefix(remaining, prefix) {
				return nil
			}
			newRemaining := remaining[len(prefix):]
			for _, branch := range branches {
				match := branch(method, newRemaining)
				if match != nil {
					return match
				}
			}
			return nil
		}
	}
}

func (compiler RequestLineCompiler[Endpoint]) Param(
	name string,
) func(branches ...RequestLineBranch[Endpoint]) RequestLineBranch[Endpoint] {
	return func(branches ...RequestLineBranch[Endpoint]) RequestLineBranch[Endpoint] {
		return func(method string, remaining string) *RequestLineMatch[Endpoint] {
			if !strings.HasPrefix(remaining, "/") {
				return nil
			}
			prefixDropped := remaining[1:]
			slashOffset := strings.IndexByte(prefixDropped, '/')
			var capture, newRemaining string
			if slashOffset == -1 {
				capture = prefixDropped
				newRemaining = ""
			} else {
				capture = prefixDropped[:slashOffset]
				newRemaining = prefixDropped[slashOffset:]
			}
			for _, branch := range branches {
				match := branch(method, newRemaining)
				if match != nil {
					return match.withParam(name, capture)
				}
			}
			return nil
		}
	}
}

func makeMethodMatcher[Endpoint any](target string, endpoint Endpoint) RequestLineBranch[Endpoint] {
	return func(method string, remaining string) *RequestLineMatch[Endpoint] {
		if remaining != "" {
			return nil
		}
		if method != target {
			return nil
		}
		return &RequestLineMatch[Endpoint]{endpoint, map[string]string{}}
	}
}

func (compiler RequestLineCompiler[Endpoint]) Get(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("GET", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Post(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("POST", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Put(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("PUT", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Delete(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("DELETE", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Options(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("OPTIONS", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Patch(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("PATCH", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Head(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("HEAD", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Connect(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("CONNECT", endpoint)
}

func (compiler RequestLineCompiler[Endpoint]) Trace(endpoint Endpoint) RequestLineBranch[Endpoint] {
	return makeMethodMatcher("TRACE", endpoint)
}
