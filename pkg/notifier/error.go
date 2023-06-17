package notifier

import "errors"

var (
	// ErrSendMessage is used when email or text message fail transmitting
	ErrSendMessage = errors.New("error_sending_message")
)