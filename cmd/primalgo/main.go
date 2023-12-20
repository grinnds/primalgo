package main

import (
	"log"
	"primalgo"
)

func main() {
	for i := uint64(1000); i < 2000; i++ {
		N, err := primalgo.NewMorrisonNumber(1, i)
		if err != nil {
			log.Fatalf("error initing morrison number: %v", N)
			return
		}

		if primalgo.MorrisonTest(N) {
			log.Printf("%v is prime", N)
		} else {
			log.Printf("%v is not prime", N)
		}
	}
}
