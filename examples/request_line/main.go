package main

import (
	"fmt"
	"github.com/unexcitingcode/http-routing"
)

func MakeRoutes[Branch any, Out any](dsl http_routing.Compiler[string, Branch, Out]) Out {
	return dsl.Root("Missing")(
		dsl.Path("/")(dsl.Get("IndexRender")),
		dsl.Path("/log_in")(
			dsl.Get("LogInRender"),
			dsl.Post("LogInProcess"),
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
	)
}

func main() {
	routes := MakeRoutes(http_routing.NewRequestLineCompiler[string]())
	request := http_routing.RequestLine{Method: "GET", Path: "/users/1337"}
	match := routes(request)
	fmt.Printf("%+v\n", match)
}
