package main

import (
	"log"
	"net/http"
	"compute-server/config"
	"fmt"
)

func main()  {
	router := NewRouter()
	configuration, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error occured while reading configuration %v\n", err)
	}
	log.Fatal(http.ListenAndServe(configuration.BindTo, router))
}
