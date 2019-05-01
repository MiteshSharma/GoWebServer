package slack

import (
	"context"

	"github.com/MiteshSharma/project/model"
	"github.com/nlopes/slack"
)

type Slack struct {
}

func New() *Slack {
	slack := &Slack{}
	return slack
}

func (s Slack) Send(to string, message model.NotificationMessage) error {
	slackClient := slack.New("Token")
	_, _, err := slackClient.PostMessageContext(
		context.TODO(),
		"channel",
		slack.MsgOptionText(message.Message, false))
	return err
}
