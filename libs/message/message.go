package message

import (
	"github.com/ika108/gsyslog/libs/input"
	"github.com/ika108/gsyslog/libs/output"
	"time"
)

type message struct {
	time.Time
	msg string
	output *output
	input *input
}