package main

import "fmt"

import "github.com/Noah-Huppert/k8s-deploy/config"

// main is the command line tool's entry point
func main() {
	// Load config
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Fatalf("error loading configuration: %s\n", err.Error())
	}

	fmt.Print(cfg)
}
