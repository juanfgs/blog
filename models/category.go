package models




type Category struct {
	Id int
	Title string `orm:"size(100)"`
	Posts []*Post `orm:"reverse(many)"`
}

func (this *Category) TableName() string{
	return "categories"
}



