package table

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ParkManagerTable = "parkManager"

//ParkManager 停车管理员信息
type ParkManager struct {
	id         bson.ObjectId `bson:"_id"`
	openID     string        //wx openID
	actualName string        //真实姓名
	tel        string        //电话
	xqID       string        //小区ID
	bind       int           // -1-解绑 0-未审核 1-绑定
}

//InsertParkManager
func InsertParkManager(db *mgo.Database, info ParkManager) interface{} {
	c := db.C(ParkManagerTable)
	count, err := c.Find(bson.M{"openid": info.openID, "xqid": info.xqID, "bind": 1}).Count()
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "在该小区已存在绑定信息"}
	}
	err = c.Insert(&info)
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

//UpdateParkManager
func UpdateParkManager(db *mgo.Database, info ParkManager) interface{} {
	c := db.C(ParkManagerTable)
	err := c.Update(bson.M{"_id": info.id}, info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}
