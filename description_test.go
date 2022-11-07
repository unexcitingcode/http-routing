package http_routing

import (
	"reflect"
	"testing"
)

func TestDescriptionCompiler(t *testing.T) {
	dsl := NewDescriptionCompiler[string]()
	var tests = []struct {
		name     string
		result   []RouteDescription[string]
		expected []RouteDescription[string]
	}{
		{
			name:     "empty routes",
			result:   dsl.Root("missing")(),
			expected: []RouteDescription[string]{},
		},
		{
			name:     "top level get",
			result:   dsl.Root("missing")(dsl.Get("GetAll")),
			expected: []RouteDescription[string]{{Method: "GET", Endpoint: "GetAll"}},
		},
		{
			name:     "top level post",
			result:   dsl.Root("missing")(dsl.Post("PostAll")),
			expected: []RouteDescription[string]{{Method: "POST", Endpoint: "PostAll"}},
		},
		{
			name:     "top level put",
			result:   dsl.Root("missing")(dsl.Put("PutAll")),
			expected: []RouteDescription[string]{{Method: "PUT", Endpoint: "PutAll"}},
		},
		{
			name:     "top level delete",
			result:   dsl.Root("missing")(dsl.Delete("DeleteAll")),
			expected: []RouteDescription[string]{{Method: "DELETE", Endpoint: "DeleteAll"}},
		},
		{
			name:     "top level options",
			result:   dsl.Root("missing")(dsl.Options("OptionsAll")),
			expected: []RouteDescription[string]{{Method: "OPTIONS", Endpoint: "OptionsAll"}},
		},
		{
			name:     "top level patch",
			result:   dsl.Root("missing")(dsl.Patch("PatchAll")),
			expected: []RouteDescription[string]{{Method: "PATCH", Endpoint: "PatchAll"}},
		},
		{
			name:     "top level head",
			result:   dsl.Root("missing")(dsl.Head("HeadAll")),
			expected: []RouteDescription[string]{{Method: "HEAD", Endpoint: "HeadAll"}},
		},
		{
			name:     "top level connect",
			result:   dsl.Root("missing")(dsl.Connect("ConnectAll")),
			expected: []RouteDescription[string]{{Method: "CONNECT", Endpoint: "ConnectAll"}},
		},
		{
			name:     "top level trace",
			result:   dsl.Root("missing")(dsl.Trace("TraceAll")),
			expected: []RouteDescription[string]{{Method: "TRACE", Endpoint: "TraceAll"}},
		},
		{
			name: "single path",
			result: dsl.Root("missing")(dsl.Path("/sign_up")(
				dsl.Get("SignUpRender"),
				dsl.Post("SignUpProcess"),
			)),
			expected: []RouteDescription[string]{
				{Method: "GET", Path: "/sign_up", Endpoint: "SignUpRender"},
				{Method: "POST", Path: "/sign_up", Endpoint: "SignUpProcess"},
			},
		},
		{
			name: "multi path",
			result: dsl.Root("missing")(
				dsl.Path("/sign_up")(
					dsl.Get("SignUpRender"),
					dsl.Post("SignUpProcess"),
				),
				dsl.Path("/log_in")(
					dsl.Get("LogInRender"),
					dsl.Post("LogInProcess"),
				),
			),
			expected: []RouteDescription[string]{
				{Method: "GET", Path: "/sign_up", Endpoint: "SignUpRender"},
				{Method: "POST", Path: "/sign_up", Endpoint: "SignUpProcess"},
				{Method: "GET", Path: "/log_in", Endpoint: "LogInRender"},
				{Method: "POST", Path: "/log_in", Endpoint: "LogInProcess"},
			},
		},
		{
			name: "single param",
			result: dsl.Root("missing")(dsl.Path("/users")(dsl.Param("user_id")(
				dsl.Get("FetchUserById"),
			))),
			expected: []RouteDescription[string]{
				{Method: "GET", Path: "/users/{user_id}", Endpoint: "FetchUserById"},
			},
		},
		{
			name: "rest resource",
			result: dsl.Root("missing")(dsl.Path("/users")(
				dsl.Post("CreateUser"),
				dsl.Param("user_id")(
					dsl.Get("FetchUserById"),
					dsl.Put("UpdateUserById"),
					dsl.Delete("DeleteUserById"),
				),
			)),
			expected: []RouteDescription[string]{
				{Method: "POST", Path: "/users", Endpoint: "CreateUser"},
				{Method: "GET", Path: "/users/{user_id}", Endpoint: "FetchUserById"},
				{Method: "PUT", Path: "/users/{user_id}", Endpoint: "UpdateUserById"},
				{Method: "DELETE", Path: "/users/{user_id}", Endpoint: "DeleteUserById"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !reflect.DeepEqual(test.result, test.expected) {
				t.Errorf("got %+v, want %+v", test.result, test.expected)
			}
		})
	}
}
