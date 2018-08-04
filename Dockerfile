FROM alpine:latest

# Directory
RUN mkdir -p /opt/k8s-deploy
WORKDIR /opt/k8s-deploy

# Install pre-requisits
RUN apk --update add git docker

# 
