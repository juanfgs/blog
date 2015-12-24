package models

import "time"

type Page struct {
	Id int
	Title string `orm:"size(100)"`
	Tagline string `orm:"size(100);null"`
	Slug string `orm:"size(100);null"`
	Content string `orm:"type(text)"`
	ContentType string `orm:"size(100)"`
	Published bool `orm:default(false)`
	Description string `orm:"type(text)"`
	Keywords string `orm:"size(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Photo string `orm:"null"`
}

func (this *Page) TableName() string {
	return "pages"
}




