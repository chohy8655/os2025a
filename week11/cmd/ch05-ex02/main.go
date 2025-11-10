package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	weights, err := datafile.GetFloats("meatweight.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for i := 0; i < len(weights); i++ {
		sum = sum + weights[i]
	}
	weeks := float64(len(weights))
	fmt.Println("평균:", sum/weeks)
}
