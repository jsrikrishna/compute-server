package main

import (
	"net/http"
	"fmt"
	"html"
	"time"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Run(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Okay I will run")
}

func Compute(w http.ResponseWriter, r *http.Request){
	computationResults := Computations{
		Computation{ProblemType:"Factorial2", Completed:true, TimeTaken:time.Now()},
		Computation{ProblemType:"Factorial2", Completed:false, TimeTaken:time.Now()},
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(computationResults); err != nil {
		panic(err)
	}
}