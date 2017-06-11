package main

import (
	"net/http"
	"fmt"
	"html"
	"time"
	"encoding/json"
	"github.com/shirou/gopsutil/mem"
)

func (computeServer *ComputeServer) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func (computeServer *ComputeServer) Run(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Okay I will run")
}

func (computeServer *ComputeServer)Compute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, i got a request !!!")
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

func (computeServer *ComputeServer) GetSystemResources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("System Resource Request")
	v, _ := mem.VirtualMemory()
	data := SystemResourcesData{
		Address: computeServer.Configuration.BindTo,
		UsedPercent: v.UsedPercent,
		Free: v.Free,
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}