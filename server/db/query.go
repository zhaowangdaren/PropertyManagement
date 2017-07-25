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
func QueryComunityInfo(db *mgo.Database, name string, pageNo int, pageSize int) interface{} {
	community := table.Community{}
	community.Name = name
	finds := table.FindCommunities(db, community, pageNo, pageSize)
	return finds
}

//QueryComunityKVs 查询社区信息
func QueryComunityKVs(db *mgo.Database, kvs map[string]interface{}) string {
	finds := table.FindCommunitiesKV(db, kvs)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryComunityDistinct 查询社区信息
func QueryComunityDistinct(db *mgo.Database, key string) interface{} {
	finds := table.FindCommunityDistincts(db, key)
	return finds
}

//QueryXiaoQuInfo 查询小区infos
func QueryXiaoQuInfo(db *mgo.Database, name string, pageNo int, pageSize int) interface{} {
	xiaoQu := table.XiaoQu{}
	xiaoQu.Name = name

	finds := table.FindXiaoQus(db, xiaoQu, pageNo, pageSize)
	return finds
}

//QueryXQDistinct 查询社区信息
func QueryXQDistinct(db *mgo.Database, key string) interface{} {

	finds := table.FindXQDistincts(db, key)
	return finds
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

//QueryHouse 查询街道用户结果集
func QueryHouse(db *mgo.Database, buildNo string, pageNo int, pageSize int) string {

	info := table.House{}
	info.BuildNo = buildNo

	finds := table.FindHouses(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}
