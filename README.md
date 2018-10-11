Project status: Complete | Maintaining

# Kubernetes Deploy
Kubernetes deployment command line tool.

# Table Of Contents
- [Overview](#overview)
- [Install](#install)
- [Usage](#usage)

# Overview
The Kubernetes deploy tool packages an opinionated Kubernetes deployment 
process into a single easy to use script.

It will complete the following steps:

- Build and tag Docker image
- Push Docker image to registry
- Deploy Helm charts to Kubernetes cluster

Benefits:

- Same deployment process everywhere
	- You do not have to replicate the same deploy code across every single 
		repository
- Can run anywhere
	- No continuous integration product specific configuration
- Easy to setup and use
	- A simple bash script
	- Just download and run

# Install
Add the Kubernetes deploy repository as a submodule:

```
git submodule add git@github.com:Noah-Huppert/k8s-deploy.git deploy/k8s-deploy
```

# Usage
Execute the `k8s-deploy` script.  
