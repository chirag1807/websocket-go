package request

import "time"

type Chat struct {
	Sender   int64     `json:"sender"`
	Receiver int64     `json:"receiver"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time,omitempty"`
}
