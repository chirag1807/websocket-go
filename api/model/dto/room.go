package dto

type RoomInfo struct {
	ID        int64   `json:"id,omitempty"`
	RoomName  string  `json:"roomname,omitempty"`
	CreatedBy int64   `json:"createdby,omitempty"`
	Members   []int64 `json:"members,omitempty"`
}
