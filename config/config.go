package config

import "flag"
import "regexp"
import "os"
import "errors"
import "fmt"

import "gopkg.in/src-d/go-git.v4"

// NewConfig loads configuration options from command line arguments.
// Defaults Config.Version to the latest git commit hash if the tool is run
// in a git repository.
func NewConfig() (*Config, error) {
	// Define flags
	var containerRepo string
	flag.StringVar(&containerRepo, "container-repo", "", "Name of repository "+
		"to push container image to")

	var containerDir string
	flag.StringVar(&containerDir, "container-dir", "", "Directory where the "+
		"container Dockerfile is located")

	var version string

	defaultVersion, err := getLatestGitCommitHash()
	if err != nil {
		return nil, fmt.Errorf("error retrieving latest git repo commit hash"+
			": %s", err.Error())
	}

	flag.StringVar(&version, "version", defaultVersion, "Release version. "+
		"If the working directory contains a git repository defaults to "+
		"the hash of the git HEAD")

	// Parse
	flag.Parse()

	// Validate argument values
	// ... container-repo
	if len(containerRepo) == 0 {
		return nil, errors.New("container-repo flag value cannot be empty")
	}

	containerRepoRegexp, err := regexp.Compile("^.*/.*$")
	if err != nil {
		return nil, fmt.Errorf("error compiling container-repo validation "+
			"regexp: %s", err.Error())
	}

	if !containerRepoRegexp.MatchString(containerRepo) {
		return nil, fmt.Errorf("container-repo flag value must be in format: "+
			"\"owner/name\", was: \"%s\"", containerRepo)
	}

	// ... container-dir
	if len(containerDir) == 0 {
		return nil, errors.New("container-dir argument value cannot be empty")
	}

	if _, err := os.Stat(containerDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("container-dir \"%s\" does not exist",
			containerDir)
	}

	// ... version
	if len(version) == 0 {
		return nil, errors.New("version flag value cannot be empty")
	}

	// Assemble config obj
	return &Config{
		ContainerRepo: containerRepo,
		ContainerDir:  containerDir,
		Version:       version,
	}, nil
}

// Config holds tool configuration options
type Config struct {
	// ContainerRepo is the name of the repository which container images will
	// be pushed to
	ContainerRepo string

	// ContainerDir is the directory which containers the container Dockerfile.
	// Defaults to the current working directory.
	ContainerDir string

	// Version is the deployment version.
	Version string
}

// getLatestGitCommitHash returns the hash of the most recent git commit of a
// repository in the working directory. If no git repository exists an empty
// string is returned.
func getLatestGitCommitHash() (string, error) {
	// Open git repo
	gitRepo, err := git.PlainOpen(".")

	// If no repo exists
	if err == git.ErrRepositoryNotExists {
		return "", nil
	}

	if err != nil {
		return "", fmt.Errorf("error reading git repository in working "+
			"directory: %s", err.Error())
	}

	// Get most recent commit
	gitHead, err := gitRepo.Head()
	if err != nil {
		return "", fmt.Errorf("error retrieving git repository head: %s",
			err.Error())
	}

	return gitHead.Hash().String(), nil
}
