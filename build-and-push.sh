#!/bin/bash

version=$1
if [ "$version" == "" ]; then
	echo "version not provided"
	exit 1
fi

## push for github actions
docker build . -t alexheld/devctl-release-bot:$version -f Dockerfile
docker push alexheld/devctl-release-bot:$version
