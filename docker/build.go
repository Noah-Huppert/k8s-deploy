package docker

import (
	"context"
	//	"errors"
	"fmt"
	//	"io/ioutil"

	"github.com/Noah-Huppert/k8s-deploy/gzip"

	//	"github.com/Noah-Huppert/golog"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

// BuildContainer builds a Docker image. With the context in the specified
// directory and tagged with the repo and version values combined by a colon.
func BuildContainer(ctx context.Context, dir, repo, version string) error {

	// Archive directory
	archiveBuffer, err := gzip.ArchiveDir(dir)
	if err != nil {
		return fmt.Errorf("error creating archive of Docker build context "+
			"directory: %s", err.Error())
	}

	// Create docker client
	dockerClient, err := docker.NewClient("unix:///var/run/docker.sock", "",
		nil, nil)
	if err != nil {
		return fmt.Errorf("error creating Docker client: %s", err.Error())
	}

	// Call docker image build
	_, err = dockerClient.ImageBuild(ctx, archiveBuffer,
		types.ImageBuildOptions{
			Tags: []string{
				fmt.Sprintf("%s:%s", repo, version),
			},
		})
	if err != nil {
		return fmt.Errorf("error building docker image: %s", err.Error())
	}

	/*
		respBytes, err := ioutil.ReadAll(buildResp.Body)
		if err != nil {
			return fmt.Errorf("error reading docker create image response: %s",
				err.Error())
		}

		if respBytes == nil || len(respBytes) == 0 {
			return errors.New("error no response from docker daemon")
		}

		respStr := string(respBytes)

		logger.Debug("resp: %#v", respStr)
	*/

	return nil
}
