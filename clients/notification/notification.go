package notification

import "fmt"

type Notifier interface {
	Notify(destination, message string) error
}

type SimpleNotifier struct{}

func NewNotifier() Notifier {
	return SimpleNotifier{}
}

func (n SimpleNotifier) Notify(destination, message string) error {
	fmt.Printf("notify %s: %s\n", destination, message)
	return nil
}
