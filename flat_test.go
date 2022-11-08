package http_routing

import (
	"reflect"
	"testing"
)

func TestFlatRouteTranspilerWithRequestLineCompiler(t *testing.T) {
	compiler := NewRequestLineCompiler[string]()
	dsl := NewFlatRouteTranspiler(compiler)
	routes := dsl.Root("Missing")(
		dsl.Get("/", "IndexRender"),
		dsl.Get("/log_in", "LogInRender"),
		dsl.Post("/log_in", "LogInProcess"),
		dsl.Get("/exhaustive", "ExhaustiveGet"),
		dsl.Post("/exhaustive", "ExhaustivePost"),
		dsl.Put("/exhaustive", "ExhaustivePut"),
		dsl.Delete("/exhaustive", "ExhaustiveDelete"),
		dsl.Options("/exhaustive", "ExhaustiveOptions"),
		dsl.Patch("/exhaustive", "ExhaustivePatch"),
		dsl.Head("/exhaustive", "ExhaustiveHead"),
		dsl.Connect("/exhaustive", "ExhaustiveConnect"),
		dsl.Trace("/exhaustive", "ExhaustiveTrace"),
		dsl.Get("/sign_up", "SignUpRender"),
		dsl.Post("/sign_up", "SignUpProcess"),
		dsl.Post("/users", "ApiCreateUser"),
		dsl.Get("/users/{user_id}", "ApiFetchUser"),
		dsl.Put("/users/{user_id}", "ApiUpdateUser"),
		dsl.Delete("/users/{user_id}", "ApiDeleteUser"),
	)
	var tests = []struct {
		name     string
		request  RequestLine
		expected RequestLineMatch[string]
	}{
		{
			name:    "match missing route",
			request: RequestLine{Method: "GET", Path: "/idk"},
			expected: RequestLineMatch[string]{
				Endpoint: "Missing",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match missing route via terminal param",
			request: RequestLine{Method: "GET", Path: "/pre_match"},
			expected: RequestLineMatch[string]{
				Endpoint: "Missing",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match missing route via incomplete param",
			request: RequestLine{Method: "GET", Path: "/pre_match/matches"},
			expected: RequestLineMatch[string]{
				Endpoint: "Missing",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match index route",
			request: RequestLine{Method: "GET", Path: "/"},
			expected: RequestLineMatch[string]{
				Endpoint: "IndexRender",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and get",
			request: RequestLine{Method: "GET", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustiveGet",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and post",
			request: RequestLine{Method: "POST", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustivePost",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and put",
			request: RequestLine{Method: "PUT", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustivePut",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and delete",
			request: RequestLine{Method: "DELETE", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustiveDelete",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and options",
			request: RequestLine{Method: "OPTIONS", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustiveOptions",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and patch",
			request: RequestLine{Method: "PATCH", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustivePatch",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and head",
			request: RequestLine{Method: "HEAD", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustiveHead",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and connect",
			request: RequestLine{Method: "CONNECT", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustiveConnect",
				Params:   map[string]string{},
			},
		},
		{
			name:    "match path and trace",
			request: RequestLine{Method: "TRACE", Path: "/exhaustive"},
			expected: RequestLineMatch[string]{
				Endpoint: "ExhaustiveTrace",
				Params:   map[string]string{},
			},
		},
		{
			name:    "rest resource create",
			request: RequestLine{Method: "POST", Path: "/users"},
			expected: RequestLineMatch[string]{
				Endpoint: "ApiCreateUser",
				Params:   map[string]string{},
			},
		},
		{
			name:    "rest resource get",
			request: RequestLine{Method: "GET", Path: "/users/1337"},
			expected: RequestLineMatch[string]{
				Endpoint: "ApiFetchUser",
				Params:   map[string]string{"user_id": "1337"},
			},
		},
		{
			name:    "rest resource put",
			request: RequestLine{Method: "PUT", Path: "/users/1337"},
			expected: RequestLineMatch[string]{
				Endpoint: "ApiUpdateUser",
				Params:   map[string]string{"user_id": "1337"},
			},
		},
		{
			name:    "rest resource delete",
			request: RequestLine{Method: "DELETE", Path: "/users/1337"},
			expected: RequestLineMatch[string]{
				Endpoint: "ApiDeleteUser",
				Params:   map[string]string{"user_id": "1337"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := routes(test.request)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %+v, want %+v", result, test.expected)
			}
		})
	}
}
