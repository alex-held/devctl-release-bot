#! /bin/bash

VERSION=v0.0.40
docker run --rm -v $(pwd):/home/app alex-held/devctl-release-bot:$VERSION devctl-release-bot template
