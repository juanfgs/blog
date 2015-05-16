package models

import(
	"time"
)


type Post struct {
	Id int
	Title string `orm:"size(100)"`
	Tagline string `orm:"size(100);null"`
	Content string `orm:"type(text)"`
	ContentType string `orm:"size(100)"`	
	Published bool `orm:default(false)`
	Description string `orm:"type(text)"`
	Keywords string `orm:"size(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Photo string `orm:"null"`
	Category *Category `orm:"rel(fk);null"`
	Author *User `orm:"rel(fk)"`
	Comments []*Comment `orm:"reverse(many)"`
}

func (this *Post) TableName() string{
	return "posts"
}



