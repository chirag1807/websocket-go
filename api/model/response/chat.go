package response

import "time"

type Chat struct {
	ID       int64     `json:"id"`
	Sender   int64     `json:"sender"`
	Receiver int64     `json:"receiver"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
}
