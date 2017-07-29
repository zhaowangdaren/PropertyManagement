package table

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Admin 管理员表格
type Admin struct {
	UserName string //用户名
	Password string //
	Level    string //权限等级
}

//AdminTableName admin table name
const AdminTableName = "Admin"

//InsertAdmin 插入user
func InsertAdmin(db *mgo.Database, user Admin) string {
	c := db.C(AdminTableName)
	count, err := c.Find(bson.M{"username": user.UserName}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return "已存在用户名为" + user.UserName + ", 请重新设置用户名"
	}
	err = c.Insert(&user)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindAdmin 查找用户
func FindAdmin(db *mgo.Database, userName string) Admin {
	c := db.C(AdminTableName)
	query := c.Find(bson.M{"username": userName})
	count, _ := query.Count()
	result := Admin{}
	if count != 0 {
		query.One(&result)
	}
	return result
}
