package main

import (
	"github.com/astaxie/beego/orm"
	"jjj/models"
	"fmt"
)

func main() {
	o := orm.NewOrm()
	if num,err := o.Delete(&models.User{Id:1});err == nil {
		fmt.Println(num)
	}
}
