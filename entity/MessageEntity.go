package entity

import "time"

type MessageStruct struct {
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MessageStruct) TableName() string {
	return "message"
}
