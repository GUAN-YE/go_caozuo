package main

import (
	"github.com/astaxie/beego/orm"
	"jjj/models"
	"fmt"
)

func main() {
	o :=orm.NewOrm()
	var user models.User
	user.Name = "slence"
	id,err := o.Insert(&user)
	if err==nil{
		fmt.Println(id)
	}
}
