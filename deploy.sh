#!/usr/bin/env bash
#
# deploy - Deploys api to Kubernetes
#
# USAGE
#	deploy.sh [OPTIONS]
#
# OPTIONS
#	--purge    Re-download k8s-deploy tool
#	--rm       Remove deploy

set -e

# Constants
docker_repo="noahhuppert/k8s-deploy"
deploy_dir="deploy"

# Parse options
while [ ! -z "$1" ]; do
	if [[ "$1" == "--purge" ]]; then
		option_purge="true"
		shift
	elif [[ "$1" == "--rm" ]]; then
		option_rm="true"
		shift
	else
		echo "Error: unknown option \"$1\"" >&2
		exit 1
	fi
done

# Download k8s-deploy
deploy_tools_dir="./deploy/tools"
k8s_deploy_path="$deploy_tools_dir/k8s-deploy"
k8s_deploy_dl_url="https://raw.githubusercontent.com/Noah-Huppert/k8s-deploy/master/k8s-deploy"

if [[ "$option_purge" == "true" ]]; then
	rm "$k8s_deploy_path"
fi

if [ ! -f "$k8s_deploy_path" ]; then
	echo "##############################"
	echo "# Installing k8s-deploy tool #"
	echo "##############################"

	if ! mkdir -p "$deploy_tools_dir"; then
		echo "Error: failed to create deploy tools directory" >&2
		exit 1
	fi

	if ! curl "$k8s_deploy_dl_url?cache_buster=$(date +%s)" > "$k8s_deploy_path"; then
		echo "Error: failed to download k8s-deploy tool" >&2
		exit 1
	fi

	if ! chmod +x "$k8s_deploy_path"; then
		echo "Error: failed to give k8s deploy tool execution permissions" >&2
		exit 1
	fi

	echo "Successfully installed k8s-deploy tool"
fi

if [[ "$option_rm" == "true" ]]; then
	echo "#######################"
	echo "# Removing deployment #"
	echo "#######################"
	
	if ! $k8s_deploy_path --deploy-dir "$deploy_dir" --rm-deploy; then
		echo "Error: failed to remove deployment" >&2
		exit 1
	fi

	echo "Successfully removed deployment"
else
	echo "#############"
	echo "# Deploying #"
	echo "#############"

	if ! $k8s_deploy_path --docker-repo "$docker_repo" --deploy-dir "$deploy_dir"; then
		echo "Error: failed to deploy" >&2
		exit 1
	fi

	echo "Successfully deployed"
fi
