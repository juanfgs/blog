package models

import(
	"time"
)


type Post struct {
	Id int
	Title string `orm:"size(100)"`
	Tagline string `orm:"size(100);null"`
	Content string `orm:"type(text)"`
	Published bool `orm:default(false)`
	CreatedAt time.Time
	UpdatedAt time.Time
	Photo string `orm:"null"`
	Category *Category `orm:"rel(fk);null"`
	User *User `orm:"rel(fk)"`
}

func (this *Post) TableName() string{
	return "posts"
}



