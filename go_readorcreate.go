package main

import (
	"github.com/astaxie/beego/orm"
	"jjj/models"
	"fmt"
)

func main() {
	o := orm.NewOrm()
	user := models.User{Name:"slene"}
	if created,id ,err :=o.ReadOrCreate(&user,"name");err ==nil{
		if created{
			fmt.Println("id:",id)
		} else {
			fmt.Println("id",id)
		}
	}
}
