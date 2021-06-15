package main

import (
	"log"

	"github.com/halverneus/https-redirect/lib/cli"
)

func main() {
	if err := cli.Execute(); nil != err {
		log.Fatalf("Error: %v\n", err)
	}
}
