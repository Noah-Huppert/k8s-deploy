package main

import (
	"context"
	"fmt"

	"github.com/Noah-Huppert/k8s-deploy/config"
	"github.com/Noah-Huppert/k8s-deploy/docker"

	"github.com/Noah-Huppert/golog"
)

// main is the command line tool's entry point
func main() {
	ctx := context.Background()

	// Configure logger
	logger := golog.NewStdLogger("k8s-deploy")

	// Load config
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// Build docker image
	dockerTagStr := fmt.Sprintf("%s:%s", cfg.ContainerRepo, cfg.Version)

	logger.Infof("Building docker image: %s", dockerTagStr)

	err = docker.BuildContainer(ctx, cfg.ContainerDir, cfg.ContainerRepo,
		cfg.Version)

	if err != nil {
		logger.Fatalf("error building docker container: %s", err.Error())
	}

	logger.Infof("Built docker image: %s", dockerTagStr)
}
