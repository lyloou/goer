package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	user := User{
		Id:4,
	}
	err := o.Read(&user)
	fmt.Println(err, user)
	if err == nil{
		user.Name = "好啊好"
		num, err := o.Update(&user)
		fmt.Println(num, err)
	}
}
func sample4() {
	orm.Debug = true
	o := orm.NewOrm()
	user := User{
		Name: "Bug",
	}
	err := o.Read(&user, "Name")
	fmt.Println(err, user)
}

func sample3() {
	orm.Debug = true
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM ip").Values(&maps)
	fmt.Println(num, err)
	fmt.Println(maps)
	user := User{Name: "Bug"}
	fmt.Println("before:", user)
	id, err := o.Insert(&user)
	fmt.Println(id, err)
	fmt.Println("after:", user)
	user2 := new(User)
	fmt.Println("before:", user2)
	id, err = o.Insert(user2)
	fmt.Println(id, err)
	fmt.Println("after:", user2)
	fmt.Println(orm.GetDB())
	fmt.Println(o.Driver().Name())
	fmt.Println(o.Driver().Type())
}

func sample2() {
	orm.Debug = true
	o := orm.NewOrm()
	var users [] User
	qt := o.QueryTable("user")
	qt.Filter("name__startswith", "Lou").
		Filter("id__lte", 6).
		Limit(2).
		All(&users)
	//One(&users)
	fmt.Println(users)
	var ips []Ip
	qs := o.QueryTable("ip")
	qs.All(&ips)
	fmt.Println(ips)
}
func sample1(o orm.Ormer) {
	user := User{Name: "LouUu5"}
	id, err := o.Insert(&user)
	fmt.Printf("NUM: %d, ERR:%v\n", id, err)
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR:%v\n", err)
	num, err := o.Delete(&u)
	fmt.Printf("NUM: %d, ERR:%v\n", num, err)
}

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

type Ip struct {
	Id int
	Ip string `orm:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/lou?charset=utf8", 30)
	orm.RegisterModel(
		new(User),
		new(Ip),
	)
	orm.RunSyncdb("default", false, true)
}
