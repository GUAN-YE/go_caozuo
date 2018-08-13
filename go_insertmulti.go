package main

import (
	"github.com/astaxie/beego/orm"
	"jjj/models"
)

func main() {
	o :=orm.NewOrm()
	var user models.User
	user := []models.User{
		{Name:"kkk"},
		{Name:"lll"}
	}
	successnums,err := o.InsertMulti(100,user)
}
