package main

import (
	"context"
	"os"

	slack "lambda/slack"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	slackURL := os.Getenv("SLACK_HOCK_URL")
	slackText := "hello world!"
	userName := "notify-test"
	iconEmoji := ":ghost:"
	channelName := os.Getenv("NOTIFY_CHANNEL")

	sl := slack.NewSlack(slackURL, slackText, userName, iconEmoji, channelName)
	sl.Send()
	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
}
