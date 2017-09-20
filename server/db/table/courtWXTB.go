package table

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CourtWXTable = "courtwx"

type CourtWX struct {
	OpenID     string
	ActualName string
	Tel        string
	UserID     string //用户表格中ID
	Bind       int    //-1-解绑 0-未审核 1-审核通过
}

func InsertCourtWX(db *mgo.Database, info CourtWX) interface{} {
	c := db.C(CourtWXTable)
	query := c.Find(bson.M{"openid": info.OpenID})
	count, err := query.Count()
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 { // 已经存在
		return gin.H{"error": 1, "data": "该微信号已绑定"}
	}
	err = c.Insert(&info)
	if err != nil {
		glog.Error(err.Error())
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}

func DelCourtWX(db *mgo.Database, openID string) error {
	c := db.C(CourtWXTable)
	err := c.Remove(bson.M{"openid": openID})
	return err
}

func FindCourtByKVs(db *mgo.Database, kvs map[string]interface{}) []CourtWX {
	querys := make(map[string]interface{})
	for k, v := range kvs {
		querys[strings.ToLower(k)] = v
	}
	c := db.C(CourtWXTable)
	query := c.Find(querys)
	var err error
	var result []CourtWX
	err = query.All(&result)
	if err != nil {
		glog.Error(err.Error())
	}
	return result
}
func FindCourtByActualName(db *mgo.Database, actualName string, pageNo int, pageSize int) interface{} {
	c := db.C(CourtWXTable)
	var result []CourtWX
	var err error
	var query *mgo.Query
	if actualName == "" {
		query = c.Find(nil)
	} else {
		query = c.Find(bson.M{"actualname": bson.M{"$regex": actualName, "$options": "$i"}})
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
	return gin.H{"error": 0, "data": gin.H{"list": result, "sum": sum}}
}

func UpdateCourtWX(db *mgo.Database, info CourtWX) interface{} {
	c := db.C(CourtWXTable)
	err := c.Update(bson.M{"openid": info.OpenID}, info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}
