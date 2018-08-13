package main

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func main() {
	o := orm.NewOrm()
	user := new(user)
	user.name = "slene"

	fmt.Println(o.Insert(user))
	user.name = "your"
	fmt.Println(o.Update(user))
	fmt.Println(o.Read(user))
	fmt.Println(o.Delete(user))
	o := orm.NewOrm()
	user :=user{Id:1}
	err := o.Read(&user)
	if err ==orm.ErrNoRows{
		fmt.Println('查不到')
	}else if err ==orm.ErrMissPK{
		fmt.Println('找不到主键')
	}else {
		fmt.Println(user.id,user.name)
	}
}