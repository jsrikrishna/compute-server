package main

import (
	"time"
	"compute-server/config"
)

type Computation struct {
	ProblemType string `json:"name"`
	Completed   bool        `json:"completed"`
	TimeTaken   time.Time `json:"due"`
}

type SystemResourcesData struct {
	Address string `json: "Address"`
	UsedPercent float64 `json: "usedPercent"`
	Free uint64 `json: "free"`
}

type Computations []Computation

type ComputeServer struct {
	Configuration config.Configuration
}