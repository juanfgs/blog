package models

import (
	"time"
)

type Comment struct {
	Id int
	Commenter string `orm:"size(100)"`
	Comment string `orm:"type(text)"`
	Post *Post `orm:"rel(fk);null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this *Comment) TableName() string{
	return "comments"
}
