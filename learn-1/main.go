//beego 学习orm task
package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"github.com/astaxie/beego/toolbox"
	"math/rand"
)

type User struct {
	Id           int
	Name         string `orm:"size(100)"`
	UserName     string  `orm:"size(100);unique"`
	PassWord     string `orm:"site(100)"`
	Test         string `orm:"null"`
	RegisterTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "test.db")
	orm.RegisterModel(new(User))
	orm.Debug = true
	orm.RunCommand()
}
func GetRandomString(len1 int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len1; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func main() {
	defer toolbox.StopTask()
	tk1 := toolbox.NewTask("tk1", "0/3 * * * * *", func() error {
		var user User
		o := orm.NewOrm()
		user.Name = "yang1"
		user.UserName = GetRandomString(12)
		o.Insert(&user)
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	select {
	}
}
