package table //社区

import (
	"fmt"
	"log"
	"strings"

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

//InsertCommunityTB 插入
func InsertCommunityTB(db *mgo.Database, comm Community) string {
	c := db.C(CommunityTableName)
	count, err := c.Find(bson.M{"name": comm.Name}).Count()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	if count > 0 {
		return "已存在街道名为" + comm.Name + ", 请重新设置街道名"
	}
	err = c.Insert(&comm)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindCommunities 查询社区信息
func FindCommunities(db *mgo.Database, community Community, pageNo int, pageSize int) Communities {
	c := db.C(CommunityTableName)
	var result Communities
	var err error
	if community == (Community{}) {
		err = c.Find(nil).All(&result.Communities)
	} else {
		err = c.Find(bson.M{"name": community.Name}).All(&result.Communities)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func FindCommunitiesKV(db *mgo.Database, kvs map[string]interface{}) Communities {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	fmt.Println("FindCommunitiesKV", query)

	c := db.C(CommunityTableName)
	var result Communities
	err := c.Find(query).All(&result.Communities)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//FindCommunityDistincts 查找key对应的不同values
func FindCommunityDistincts(db *mgo.Database, key string) []string {
	c := db.C(CommunityTableName)
	var result []string
	var err error
	err = c.Find(nil).Distinct(key, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
