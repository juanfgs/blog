package models

import(
	"time"
)

const (
	NotificationStatusNew = iota
	NotificationStatusRead
	NotificationStatusDismissed
)

const(
	NotificationTypeComment = iota
	NotificationTypeContact
)

type Notification struct {
	Id int
	Message string `orm:"size(255)"`
	Type int64
	Status int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this *Notification) TableName() string {
	return "notifications"
}

