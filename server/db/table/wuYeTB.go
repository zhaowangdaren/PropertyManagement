package table

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//WuYeTableName table name
const WuYeTableName = "WuYe"

//WuYe 物业信息
type WuYe struct {
	Name                 string //物业公司名称
	Street               string //所在街道
	Community            string //所在社区
	XiaoQu               string //所在小区
	LegalPerson          string //独立法人
	Tel                  string //
	WuYeZiZhi            string //物业资质
	WuYeXinZhi           string //物业性质
	XiaoQuEnv            string //小区环境
	XiaoQuCleaning       string //小区保洁
	GreenVegetatio       string //绿化植被
	GuanYangBaoHu        string //管养保护
	XiaoFanJianCha       string //消防检查
	DianTiBaoYang        string //电梯保养
	CarParkInOrder       string //车辆有序停放
	YeZhuCommunity       string //业主委员会
	YeZhuCommunityTel    string //业主委员会联系方式
	PastYearLevel        string //上一年物业等级
	JianYiZhengGaiCuoShi string //建议整改措施
	Img1                 string
	Img2                 string
	Img3                 string
	Img4                 string
	Img5                 string
}

//WuYes 集合
type WuYes struct {
	WuYes []WuYe
}

//InsertWuYe 插入user
func InsertWuYe(db *mgo.Database, info WuYe) string {
	c := db.C("WuYe")
	count, err := c.Find(bson.M{"name": info.Name}).Count()
	if err != nil { //查询出错或记录不存在
		log.Fatal(err)
	}
	if count > 0 {
		return "已存在物业公司：" + info.Name + ", 请重新设置公司名"
	}
	err = c.Insert(&info)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return Succ
}

//FindWuYes 查询小区信息集
func FindWuYes(db *mgo.Database, name string, pageNo int, pageSize int) interface{} {
	c := db.C(WuYeTableName)
	var result []WuYe
	var err error
	if name == "" {
		err = c.Find(nil).All(&result)
	} else {
		err = c.Find(bson.M{"name": name}).All(&result)
	}
	if err != nil {
		log.Println(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	return gin.H{"error": 0, "data": result}
}

//FindWuYeKVs 通过一系列的K-V来查找记录
func FindWuYeKVs(db *mgo.Database, kvs map[string]interface{}) interface{} {
	query := make(map[string]interface{})
	for k, v := range kvs {
		query[strings.ToLower(k)] = v
	}
	c := db.C(WuYeTableName)
	var result []WuYe
	err := c.Find(query).All(&result)
	if err != nil {
		log.Fatal(err)
		return gin.H{"error": 1, "data": err.Error()}
	}
	if result == nil {
		return gin.H{"error": 1, "data": "没有查询到结果"}
	}
	return gin.H{"error": 0, "data": result}
}
