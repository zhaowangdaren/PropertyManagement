package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PMUser 绑定的微信用户，包括正在申请绑定
type PMUser struct {
	PMID       string // 物业公司ID
	ActualName string // 真实姓名
	OpenID     string // OpenId
	Tel        string // 电话号码
	Bind       int    // 0-，1-绑定成功
}

//PMUsers 微信物业公司用户集
type PMUsers struct {
	PMUsers []PMUser
}

//PMUserTableName table name
const PMUserTableName = "PMUser"

//InsertPMUser 插入user
func InsertPMUser(db *mgo.Database, info PMUser) interface{} {
	c := db.C(PMUserTableName)
	count, err := c.Find(bson.M{"openid": info.OpenID, "pmid": info.PMID}).Count()
	if err != nil { //查询出错或记录不存在
		glog.Warning("查询数量" + err.Error())
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "微信号已申请了绑定该物业公司"}
	}
	err = c.Insert(&info)
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": ""}
}
func UpdatePMUser(db *mgo.Database, info PMUser) interface{} {
	c := db.C(PMUserTableName)
	err := c.Update(bson.M{"openid": info.OpenID, "pmid": info.PMID}, info)
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindPMUsers 查询小区信息集
func FindPMUsers(db *mgo.Database, actualName string, pageNo int, pageSize int) interface{} {
	c := db.C(PMUserTableName)
	var query *mgo.Query
	if actualName == "" {
		query = c.Find(nil)
	} else {
		query = c.Find(bson.M{"actualname": bson.M{"$regex": actualName, "$options": "$i"}})
	}
	sum, _ := query.Count()
	var err error
	var result []PMUser

	if pageSize != 0 {
		err = query.Skip(pageNo * pageSize).Limit(pageSize).All(&result)
	} else { //查询所有
		err = query.All(&result)
	}
	if err != nil {
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": gin.H{"pmUsers": result, "sum": sum}}
}

func FindPMUserByKV(db *mgo.Database, key string, value string) PMUser {
	c := db.C(PMUserTableName)
	var result PMUser
	c.Find(bson.M{strings.ToLower(key): value}).One(&result)
	return result
}

func FindPMUserByKVs(db *mgo.Database, kvs map[string]interface{}) []PMUser {
	querys := make(map[string]interface{})
	for k, v := range kvs {
		querys[strings.ToLower(k)] = v
	}
	c := db.C(PMUserTableName)
	query := c.Find(querys)
	var err error
	var result []PMUser
	err = query.All(&result)
	if err != nil {
		glog.Error(err.Error())
	}
	return result
}

func FindDistinctPMUserByKVs(db *mgo.Database, kvs map[string]interface{}) []PMUser {
	querys := make(map[string]interface{})
	for k, v := range kvs {
		querys[strings.ToLower(k)] = v
	}
	c := db.C(PMUserTableName)
	query := c.Find(querys)
	var err error
	var result []PMUser
	err = query.Distinct("openid", &result)
	if err != nil {
		glog.Error(err.Error())
	}
	return result
}

//FindPMUserDistinct 筛选key对应的value
func FindPMUserDistinct(db *mgo.Database, key string) interface{} {
	c := db.C(PMUserTableName)
	var result []string
	var err error
	err = c.Find(nil).Distinct(strings.ToLower(key), &result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}
