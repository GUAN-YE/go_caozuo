package main

import (
	"github.com/astaxie/beego/orm"
	"jjj/models"
	"fmt"
	"debug/dwarf"
)

func main() {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs.Filter("id",1)
	qs.Filter("profile__age",18)  //
	qs.Filter("profile__age__gt",18) // __gt 大于
	qs.Filter("profile_age_gte",18)  //  __gte 大于等于
	qs.Filter("profile__age__in",18,20).Exclude("profile__age__it",1000)  //  __in  在范围之内，  __it 小于
	qs.Filter("name","lll")  // 查找 name是lll的数据
	qs.Filter("name__exact",slene)   //   exact 是等于的意思
	qs.Filter("profile_id",nil)
	qs.Filter("name__iexact","slene")  //  __iexact  与mysql 的like 作用一样
	cnt,err := o.QueryTable("user").Count() // count 返回结果行数
	fmt.Println("jjj",cnt,err)
	exist := o.QueryTable("user").Filter("username","name").Exist()
	fmt.Println("%s",exist)
	num,err := o.QueryTable("user").Filter("name","slene").Update(orm.Params{"name":"astaxie"})
	fmt.Println("affected num:%s,%s",num,err)
	num,err := o.QueryTable("user").Filter("name","slene").Delete()
	fmt.Println("affected num:%s,%s",num,err)
	var users []*models.User
	qs :=o.QueryTable("user")
	i,_ :=qs.PrepareInsert()
	for _,user :=range users{
		id,err := i.Insert(user)
		if err == nil{

		}
	}
	var  users []*models.User
	num,err := o.QueryTable("user").Filter("name","slence").All(&users)
	fmt.Println("return rows num :%s,%s",num,err)

	var  user models.User
	err := o.QueryTable("user").Filter("name","slene").One(&user)
	if err == orm.ErrMultiRows {
		fmt.Println("return multi rows not one")
	}
	if err == orm.ErrNoRows{
		fmt.Println("not row found")
	}
	var maps []orm.Params
	num,err:=o.QueryTable("user").Values(&maps)
	if err ==nil {
		fmt.Println("result nums:%d\n",num)
		for _,m := range maps {
			fmt.Println(m["id"],m["name"])
		}
	}
	o := orm.NewOrm()
	var r rawster
	r = o.Raw("update user set name = ? where name = ?", "testing","slene")



}