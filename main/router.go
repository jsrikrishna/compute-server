package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(computeServer *ComputeServer) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range computeServer.GetRoutes() {
		var handler http.Handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}
