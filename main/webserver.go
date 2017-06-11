package main

import (
	"log"
	"net/http"
	"compute-server/config"
	"fmt"
)

func main()  {
	configuration, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error occured while reading configuration %v\n", err)
	}
	computeServer := &ComputeServer{
		Configuration: configuration,
	}
	router := NewRouter(computeServer)

	log.Fatal(http.ListenAndServe(configuration.BindTo, router))
}
