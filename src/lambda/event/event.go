package event

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type snsMessage struct {
	Status   string `json:"status"`
}

type slackInfo struct {
	Text string
	Username  string
	IconEmoji string
}

func GetDeployStatus(snsEvent events.SNSEvent) string {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		var message snsMessage
		if err := json.Unmarshal([]byte(snsRecord.Message), &message); err != nil {
			panic(err)
		}
		return message.Status
	}
	return ""
}

func GetSlackInfoByDeployStatus(status string) *slackInfo {
	if status == "FAILED" {
		return &slackInfo {
			Text: "デプロイに失敗しました",
			Username: "notify-deploy-fail",
			IconEmoji: ":ghost:"}
	} else if status == "CREATED" {
		return &slackInfo {
			Text: "デプロイを開始しました",
			Username: "notify-deploy-start",
			IconEmoji: ":star:"}
	}
	// status == "SUCCEEDED"
	return &slackInfo {
		Text: "デプロイに成功しました",
		Username: "notify-deploy-success",
		IconEmoji: ":clap:"}
}
