package notification

import (
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/notification/app/notification/awsSes"
	"github.com/MiteshSharma/project/notification/app/notification/slack"
	"github.com/MiteshSharma/project/notification/app/notification/smtp"
)

func GetNotification(notificationType model.NotificationType) Notification {
	var notification Notification
	switch notificationType {
	case model.SMTP:
		notification = smtp.New()
		break
	case model.AWS_SES:
		notification = awsSes.New()
		break
	case model.SLACK:
		notification = slack.New()
		break
	default:
	}
	return notification
}
