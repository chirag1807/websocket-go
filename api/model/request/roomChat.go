package request

import "time"

type RoomChat struct {
	RoomID  int64     `json:"roomid"`
	Sender  int64     `json:"sender"`
	Message string    `json:"message"`
	Time    time.Time `json:"time,omitempty"`
}

type RoomInfo struct {
	ID        int64   `json:"id,omitempty"`
	RoomName  string  `json:"roomname"`
	CreatedBy int64   `json:"createdby"`
	Members   []int64 `json:"members"`
}
