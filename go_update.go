package main

import (
	"github.com/astaxie/beego/orm"
	"jjj/models"
	"fmt"
)

func main() {
	o :=orm.NewOrm()
	user :=models.User{Id:2}
	if o.Read(&user) ==nil{
		user.Name = "myname"
		if num,err :=orm.update(&user);err ==nil{
			fmt.Println(num)
		}
	}
}
