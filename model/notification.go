package model

type NotificationData struct {
	To      string
	Type    NotificationType
	Message NotificationMessage
}

func NewNotificationData(to string, message NotificationMessage, nType NotificationType) NotificationData {
	notificationData := NotificationData{
		To:      to,
		Message: message,
		Type:    nType,
	}
	return notificationData
}

type NotificationMessage struct {
	Title   string
	Message string
	Type    string
}

func NewNotificationMessage(title string, message string, nType string) NotificationMessage {
	notificationMessage := NotificationMessage{
		Title:   title,
		Message: message,
		Type:    nType,
	}
	return notificationMessage
}

type NotificationType int

const (
	SMTP NotificationType = iota
	AWS_SES
	SLACK
)

func (n NotificationType) String() string {
	return [...]string{"SMTP", "AWS_SES", "SLACK"}[n]
}
