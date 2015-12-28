package models

type Setting struct {
	Id int
	Key string `orm:"size(100)"`
	Value string `orm:"size(255)"`

}

func (this *Setting) TableName() string{
	return "settings"
}



