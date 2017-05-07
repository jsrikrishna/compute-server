package main

import (
	"time"
)

type Computation struct {
	ProblemType string `json:"name"`
	Completed   bool        `json:"completed"`
	TimeTaken   time.Time `json:"due"`
}

type Computations []Computation