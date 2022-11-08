package main

import (
	"fmt"
	"github.com/unexcitingcode/http-routing"
)

func MakeRoutes[Branch any, Out any](compiler http_routing.Compiler[string, Branch, Out]) Out {
	dsl := http_routing.NewFlatRouteTranspiler(compiler)
	return dsl.Root("Missing")(
		dsl.Get("/", "IndexRender"),
		dsl.Get("/log_in", "LogInRender"),
		dsl.Post("/log_in", "LogInProcess"),
		dsl.Get("/sign_up", "SignUpRender"),
		dsl.Post("/sign_up", "SignUpProcess"),
		dsl.Post("/users", "ApiCreateUser"),
		dsl.Get("/users/{user_id}", "ApiFetchUser"),
		dsl.Put("/users/{user_id}", "ApiUpdateUser"),
		dsl.Delete("/users/{user_id}", "ApiDeleteUser"),
	)
}

func main() {
	routes := MakeRoutes(http_routing.NewRequestLineCompiler[string]())
	request := http_routing.RequestLine{Method: "GET", Path: "/users/1337"}
	match := routes(request)
	fmt.Printf("%+v\n", match)
}
