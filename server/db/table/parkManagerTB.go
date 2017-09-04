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
	ID         bson.ObjectId `bson:"_id"`
	OpenID     string        //wx openID
	ActualName string        //真实姓名
	Tel        string        //电话
	XQID       string        //小区ID
	Bind       int           // -1-解绑 0-未审核 1-审核通过
}

//InsertParkManager
func InsertParkManager(db *mgo.Database, info ParkManager) interface{} {
	c := db.C(ParkManagerTable)
	query := c.Find(bson.M{"openid": info.OpenID, "xqid": info.XQID})
	count, err := query.Count()
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 { // 已经存在
		return gin.H{"error": 1, "data": "在该小区已存在绑定信息"}
	}
	info.ID = bson.NewObjectId()
	err = c.Insert(&info)
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}
func FindPrakManagerByActualName(db *mgo.Database, actualName string, pageNo int, pageSize int) interface{} {
	c := db.C(ParkManagerTable)
	var result []ParkManager
	var err error
	var query *mgo.Query
	if actualName == "" {
		query = c.Find(nil)
	} else {
		query = c.Find(bson.M{"actualname": actualName})
	}

	sum, _ := query.Count()
	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"parkManagers": result, "sum": sum}}
}

//UpdateParkManager
func UpdateParkManager(db *mgo.Database, info ParkManager) interface{} {
	c := db.C(ParkManagerTable)
	err := c.Update(bson.M{"_id": info.ID}, info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

func FindParjManagerByOpenID(db *mgo.Database, openID string) interface{} {
	c := db.C(ParkManagerTable)
	var result []ParkManager
	err := c.Find(bson.M{"openid": openID}).All(&result)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
