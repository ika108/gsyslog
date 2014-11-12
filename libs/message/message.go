package message

import "time"

type message struct {
	time.Time
	msg string
}