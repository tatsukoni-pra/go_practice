package main

import (
	"context"

	slack "deploy_notify/slack"
	event "deploy_notify/event"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) (string, error) {
	slackInfoByDeployStatus := event.GetSlackInfoByDeployStatus(event.GetDeployStatus(snsEvent))
	sl := slack.NewSlack(slackInfoByDeployStatus.Text, slackInfoByDeployStatus.Username, slackInfoByDeployStatus.IconEmoji)
	sl.Send()
	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
}
