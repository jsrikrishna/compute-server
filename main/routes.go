package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
func (computeServer *ComputeServer) GetRoutes() Routes {
	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			computeServer.Index,
		},
		Route{
			"Run",
			"GET",
			"/run",
			computeServer.Run,
		},
		Route{
			"Resources",
			"POST",
			"/resources",
			computeServer.Compute,
		},
		Route{
			"SystemResources",
			"GET",
			"/systemResources",
			computeServer.GetSystemResources,
		},
	}
	return routes

}