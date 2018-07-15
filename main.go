package main

import "log"

import "github.com/Noah-Huppert/k8s-deploy/config"

// main is the command line tool's entry point
func main() {
	// Load config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading configuration: %s\n", err.Error())
	}

	log.Printf("%#v\n", cfg)
}
