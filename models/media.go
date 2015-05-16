package models

import (
	"time"
)

type Media struct {
	Id int
	Filename string `orm: "size(200)"`
	Type string `orm: "size(100)"`	
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this Media) TableName() string {
	return "media"
}
