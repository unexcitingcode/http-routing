package http_routing

import (
	"reflect"
	"testing"
)

func TestRequestLineCompiler(t *testing.T) {
	dsl := NewRequestLineCompiler[string]()
	routes := dsl.Root("Missing")(
		dsl.Path("/")(dsl.Get("IndexRender")),
		dsl.Path("/log_in")(
			dsl.Get("LogInRender"),
			dsl.Post("LogInProcess"),
		),
		dsl.Path("/exhaustive")(
			dsl.Get("ExhaustiveGet"),
			dsl.Post("ExhaustivePost"),
			dsl.Put("ExhaustivePut"),
			dsl.Delete("ExhaustiveDelete"),
			dsl.Options("ExhaustiveOptions"),
			dsl.Patch("ExhaustivePatch"),
			dsl.Head("ExhaustiveHead"),
			dsl.Connect("ExhaustiveConnect"),
			dsl.Trace("ExhaustiveTrace"),
		),
		dsl.Path("/sign_up")(
			dsl.Get("SignUpRender"),
			dsl.Post("SignUpProcess"),
		),
		dsl.Path("/users")(
			dsl.Post("ApiCreateUser"),
			dsl.Param("user_id")(
				dsl.Get("ApiFetchUser"),
				dsl.Put("ApiUpdateUser"),
				dsl.Delete("ApiDeleteUser"),
			),
		),
		dsl.Path("/pre_match")(
			dsl.Param("first")(
				dsl.Param("second")(
					dsl.Path("/post_match")(
						dsl.Get("ParamMatch"),
					),
				),
			),
		),
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
			name:    "match multi param",
			request: RequestLine{Method: "GET", Path: "/pre_match/a/b/post_match"},
			expected: RequestLineMatch[string]{
				Endpoint: "ParamMatch",
				Params:   map[string]string{"first": "a", "second": "b"},
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
