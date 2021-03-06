package models

import (
	"regexp"
	"strings"
	"time"
	"fmt"
	"log"
)

type Post struct {
	Id          int
	Title       string `orm:"size(100)"`
	Tagline     string `orm:"size(100);null"`
	Slug        string `orm:"size(100);null"`
	Content     string `orm:"type(text)"`
	ContentType string `orm:"size(100)"`
	Published   bool   `orm:default(false)`
	Description string `orm:"type(text)"`
	Keywords    string `orm:"size(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Photo       string     `orm:"null"`
	Category    *Category  `orm:"rel(fk);null"`
	Author      *User      `orm:"rel(fk)"`
	Comments    []*Comment `orm:"reverse(many)"`
}

func (this *Post) TableName() string {
	return "posts"
}

func GenerateSlug(title string) []byte {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	safe := reg.ReplaceAllString(title, "-")
	return []byte(strings.ToLower(strings.Trim(safe, "-")))

}

func (this *Post) Url() string{
	if len(this.Slug) > 0 {
		return fmt.Sprintf("/post/%s", this.Slug)
	} else {
		return fmt.Sprintf("/post/%d", this.Id )	
	}
}
