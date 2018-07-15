package config

import "flag"
import "regexp"
import "os"

import "gopkg.in/src-d/go-git.v4"

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

// NewConfig loads configuration options from command line arguments.
// Defaults Config.Version to the latest git commit hash if the tool is run
// in a git repository.
func NewConfig() (*Config, error) {
	// Container repository argument
	containerRepo := flag.String("container-repo", "", "Name of repository "+
		"to push container image to")

	// ... Validate
	if len(*containerRepo) == 0 {
		return nil, errors.New("container-repo flag value cannot be empty")
	}

	containerRepoRegexp, err := regexp.Compile("^.*/.*$")
	if err != nil {
		return nil, errors.New("error compiling container-repo validation regexp: %s",
			err.Error())
	}

	if !containerRepoRegexp.MatchString(containerRepo) {
		return nil, errors.Newf("container-repo flag value must be in format: "+
			"\"owner/name\", was: \"%s\"", containerRepo)
	}

	// Container directory argument
	containerDir := flag.String("container-dir", "", "Directory which "+
		"contains container Dockerfile")

	// ,,, Validate
	if len(containerDir) == 0 {
		return nil, errors.New("container-dir argument value cannot be empty")
	}

	if _, err := os.Stat(containerDir); os.IsNotExist(err) {
		return nil, errors.Newf("container-dir \"%s\" does not exist",
			containerDir)
	}

	// Version argument
	// ... Get most recent git commit hash for version default
	gitRepo, err := git.PlainOpen(".")
	if err != nil && err != git.ErrRepositoryNotExists {
		return nil, errors.Newf("error reading git repository in working directory: %s",
			err.Error())
	}

	gitHead, err := gitRepo.Head()
	if err != nil {
		return nil, errors.Newf("error retrieving git repository head: %s", err.Error())
	}

	defaultVersion := string(gitHead.Hash())

	// ... Get arg
	version := flag.String("version", defaultVersion, "Release version, used"+
		" as the container image tag and Helm chart app version. Defaults "+
		"to the hash of the most recent commit if invoked in a git repository")

	// ... Validate
	if len(version) == 0 {
		return nil, errors.New("version flag value cannot be empty")
	}

	// Assemble object
	return &Config{
		ContainerRepo: containerRepo,
		ContainerDir:  containerDir,
		Version:       version,
	}, nil
}
