package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

var routes = Routes{
	Route{Name: "Index", Method: GET, Pattern: "/", HandlerFunc: Index},
	Route{Name: "TodoIndex", Method: GET, Pattern: "/todos", HandlerFunc: TodoIndex},
	Route{Name: "TodoShow", Method: GET, Pattern: "/todos/{todoId}", HandlerFunc: TodoShow},
	Route{Name: "TodoCreate", Method: POST, Pattern: "/todos", HandlerFunc: TodoCreate},
}
