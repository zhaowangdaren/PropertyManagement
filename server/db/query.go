package db

import (
	"encoding/json"

	"./table"
	mgo "gopkg.in/mgo.v2"
)

//QueryStreetInfo 查询街道信息
func QueryStreetInfo(db *mgo.Database, name string, pageNo int, pageSize int) interface{} {
	street := table.Street{}
	street.Name = name
	finds := table.FindStreets(db, street, pageNo, pageSize)
	return finds
}

//QueryComunityInfo 查询社区信息
func QueryComunityInfo(db *mgo.Database, name string, pageNo int, pageSize int) string {
	community := table.Community{}
	community.Name = name
	finds := table.FindCommunities(db, community, pageNo, pageSize)
	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryComunityKVs 查询社区信息
func QueryComunityKVs(db *mgo.Database, kvs map[string]interface{}) string {
	finds := table.FindCommunitiesKV(db, kvs)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryComunityDistinct 查询社区信息
func QueryComunityDistinct(db *mgo.Database, key string) string {
	finds := table.FindCommunityDistincts(db, key)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryXiaoQuInfo 查询小区infos
func QueryXiaoQuInfo(db *mgo.Database, name string, pageNo int, pageSize int) string {
	xiaoQu := table.XiaoQu{}
	xiaoQu.Name = name

	finds := table.FindXiaoQus(db, xiaoQu, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryXQDistinct 查询社区信息
func QueryXQDistinct(db *mgo.Database, key string) string {

	finds := table.FindXQDistincts(db, key)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryXQKVs 查询
func QueryXQKVs(db *mgo.Database, kvs map[string]interface{}) string {
	finds := table.FindXQKVs(db, kvs)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryGovInfo 查询小区infos
func QueryGovInfo(db *mgo.Database, name string, pageNo int, pageSize int) string {
	gov := table.Gov{}
	gov.UserName = name

	finds := table.FindGovs(db, gov, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryStreetUserInfo 查询街道用户结果集
func QueryStreetUserInfo(db *mgo.Database, name string, pageNo int, pageSize int) string {
	info := table.StreetUser{}
	info.UserName = name

	finds := table.FindStreetUsers(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryWXUser 查询街道用户结果集
func QueryWXUser(db *mgo.Database, openID string, pageNo int, pageSize int) string {
	info := table.WXUser{}
	info.OpenID = openID

	finds := table.FindWXUsers(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryWuYe 查询街道用户结果集
func QueryWuYe(db *mgo.Database, xiaoQu string, pageNo int, pageSize int) string {

	info := table.WuYe{}
	info.XiaoQu = xiaoQu

	finds := table.FindWuYes(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryHouse 查询街道用户结果集
func QueryHouse(db *mgo.Database, buildNo string, pageNo int, pageSize int) string {

	info := table.House{}
	info.BuildNo = buildNo

	finds := table.FindHouses(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}
