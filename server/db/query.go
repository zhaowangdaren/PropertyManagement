package db

import (
	"encoding/json"

	"./table"
	mgo "gopkg.in/mgo.v2"
)

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
