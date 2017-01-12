package models

type Link struct {
	Id int
	Title string `orm:"size(100)"`
	Url string `orm:"size(255)"`
	
}

func (this *Link) TableName() string{
	return "links"
}



