package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"strings"
)

type SlackMessage struct {
	Title     string `json:"pretext,omitempty"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Username  string `json:"username,omitempty"`
	Color     string `json:"color,omitempty"`
}

type SlackMessageBuilder struct {
	sm SlackMessage
}

func NewSlackMessageBuilder() *SlackMessageBuilder {
	return &SlackMessageBuilder{}
}

func (smb *SlackMessageBuilder) SetTitle(title string) *SlackMessageBuilder {
	smb.sm.Title = title
	return smb
}

func (smb *SlackMessageBuilder) SetText(text string) *SlackMessageBuilder {
	smb.sm.Text = text
	return smb
}

func (smb *SlackMessageBuilder) SetIconEmoji(emoji string) *SlackMessageBuilder {
	smb.sm.IconEmoji = emoji
	return smb
}

func (smb *SlackMessageBuilder) SetUserName(userName string) *SlackMessageBuilder {
	smb.sm.Username = userName
	return smb
}

func (smb *SlackMessageBuilder) SetColor(color string) *SlackMessageBuilder {
	smb.sm.Color = color
	return smb
}

func (smb *SlackMessageBuilder) SendSlack() {
	const webhookUrl = "YOUR_SLACK_WEBHOOK_URL"

	pbytes, _ := json.Marshal(smb.sm)
	buff := bytes.NewBuffer(pbytes)

	http.Post(webhookUrl, "application/json", buff)
}

func Bold(s string) string {
	return "*" + s + "*"
}

func StatusToColor(subject string) string {
	if strings.Contains(subject, "unable to contact") || strings.Contains(subject, "not available") {
		return "red"
	}
	if strings.Contains(subject, "can't access") {
		return "#808080"
	}
	if strings.Contains(subject, "status of Severe") {
		return "#FFFF00"
	}
	return "good"
}

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		newMsg := NewSlackMessageBuilder()
		newMsg.SetTitle(Bold(snsRecord.Subject)).
			SetText(snsRecord.Message).
			SetIconEmoji(":robot_face:").
			SetUserName("모니터링 봇이 알립니다").
			SetColor(StatusToColor(snsRecord.Subject)).
			SendSlack()
	}
}

func main() {
	lambda.Start(HandleRequest)
}
