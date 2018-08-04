# Kubernetes Deploy
Kubernetes deployment command line tool.

# Table Of Contents
- [Overview](#overview)
- [Usage](#usage)

# Overview
The Kubernetes deploy tool packages an opinionated Kubernetes deployment 
process into a single easy to use binary.

It will complete the following steps:

- Build and tag container
- Push container to registry
- Deploy Helm chart to Kubernetes cluster

Benefits:

- Same deployment process everywhere
	- You do not have to replicate the same deploy code across every single 
		repository
- Can run anywhere
	- No continuous integration product specific configuration
- Easy to setup and use
	- A simple bash script
	- Just download and run

# Usage
Execute the `k8s-deploy` script.
