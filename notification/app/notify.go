package app

import (
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/notification/app/notification"
)

func (a *App) SendNotification(notificationData model.NotificationData) {
	notification := notification.GetNotification(notificationData.Type)
	if notification != nil {
		err := notification.Send(notificationData.To, notificationData.Message)
		if err != nil {
			a.Log.Warn("Notification sent failed", logger.Error(err))
		} else {
			a.Log.Debug("Notification is sent", logger.String("type", notificationData.Type.String()))
		}
	} else {
		a.Log.Error("Notification client not set for this type",
			logger.String("type", notificationData.Type.String()))
	}
}
