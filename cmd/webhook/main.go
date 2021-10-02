package main

import (
	"os"

	"github.com/alex-held/devctl-release-bot/pkg/releaser"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	ghToken := os.Getenv("GH_TOKEN")
	releaser := releaser.New(ghToken)

	lambda.Start(releaser.HandleActionLambdaWebhook)
}
