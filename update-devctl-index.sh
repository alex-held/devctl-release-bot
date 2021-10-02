#!/bin/bash

export DEVCTL_RELEASE_BOT_VERSION=v0.0.40

curl -LO https://github.com/alex-held/devctl-release-bot/releases/download/${DEVCTL_RELEASE_BOT_VERSION}/devctl-release-bot_${DEVCTL_RELEASE_BOT_VERSION}_linux_amd64.tar.gz
tar -xvf devctl-release-bot_${DEVCTL_RELEASE_BOT_VERSION}_linux_amd64.tar.gz
./devctl-release-bot action
