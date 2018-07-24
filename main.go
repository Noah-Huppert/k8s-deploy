package main

import (
	"github.com/Noah-Huppert/k8s-deploy/config"

	"github.com/Noah-Huppert/golog"
)

// main is the command line tool's entry point
func main() {
	// Configure logger
	logger := golog.NewStdLogger("k8s-deploy")

	// Load config
	// TODO: Find better logging library, maybe switch back to std
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	logger.Debugf("%#v", cfg)
}
