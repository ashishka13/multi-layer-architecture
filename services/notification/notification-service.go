package notification

import "multi-layer-architecture/clients/notification"

type NotificationServiceInterface interface {
	Send(destination, message string) error
}

type NotificationService struct {
	notifier notification.Notifier
}

func NewNotificationService(n notification.Notifier) NotificationServiceInterface {
	return NotificationService{notifier: n}
}

func (s NotificationService) Send(destination, message string) error {
	return s.notifier.Notify(destination, message)
}
