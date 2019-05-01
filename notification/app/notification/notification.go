package notification

import "github.com/MiteshSharma/project/model"

type Notification interface {
	Send(to string, message model.NotificationMessage) error
}
