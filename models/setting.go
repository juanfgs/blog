package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)
type Setting struct {
	Id int
	Key string `orm:"size(100)"`
	Value string `orm:"size(255)"`

}

func (this *Setting) TableName() string{
	return "settings"
}

func GetSettingsMap() map[string]string {
	var settings []Setting
	var settingsMap  map[string]string
	settingsMap = make(map[string]string)

	o := orm.NewOrm()
	o.QueryTable("settings").All(&settings)

	for _,setting := range settings {
		settingsMap[setting.Key] = setting.Value
	}
	log.Println("here")
	return settingsMap
}

func GetSetting(key string)  (Setting,error) {
	var s Setting
	o := orm.NewOrm()
	err := o.QueryTable("settings").Filter("key", key).One(&s)

	return s, err
}

