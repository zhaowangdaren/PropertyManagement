package table //社区

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//CommunityTableName 社区表格名称
const CommunityTableName = "Community"

//Community 社区
type Community struct {
	Name           string //社区名
	PersonInCharge string //负责人
	Tel            string //社区联系电话
	StreetName     string //所属街道
	Intro          string //介绍
}

//Communities 社区集
type Communities struct {
	Communities []Community
}

//InsertCommunity 插入
func InsertCommunity(db *mgo.Database, comm Community) interface{} {
	c := db.C(CommunityTableName)
	count, err := c.Find(bson.M{"name": comm.Name}).Count()
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if count > 0 {
		return gin.H{"error": 1, "data": "已存在community" + comm.Name + ", 请重新设置name"}
	}
	err = c.Insert(&comm)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": Succ}
}

//FindCommunities 查询社区信息
func FindCommunities(db *mgo.Database, community Community, pageNo int, pageSize int) interface{} {
	c := db.C(CommunityTableName)
	var result []Community
	var err error
	if community == (Community{}) {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"name": community.Name}).All(&result)
	}
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

func FindCommunitiesKV(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(CommunityTableName)
	var result []Community
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindCommunityDistincts 查找key对应的不同values
func FindCommunityDistincts(db *mgo.Database, key string) interface{} {
	c := db.C(CommunityTableName)
	var result []string
	var err error
	err = c.Find(nil).Distinct(strings.ToLower(key), &result)
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindCommunityByStreetName 查询社区信息
func FindCommunityByStreetName(db *mgo.Database, streetName string) interface{} {
	c := db.C(CommunityTableName)
	var result []Community
	err := c.Find(bson.M{"streetname": streetName}).All(&result)
	if err != nil {
		log.Print(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//DelCommunities 根据名称删除记录
func DelCommunities(db *mgo.Database, names []string) interface{} {
	c := db.C(CommunityTableName)
	var err error
	result := ""
	for _, v := range names {
		err = c.Remove(bson.M{"name": v})
		if err != nil {
			log.Println(err.Error())
			result += err.Error()
			err = nil
		}
	}
	if result != "" {
		return gin.H{"error": 1, "data": result}
	}
	return gin.H{"error": 0, "data": Succ}
}
