package db

import (
	"encoding/json"

	"./table"
	mgo "gopkg.in/mgo.v2"
)

//QueryStreetInfo 查询街道信息
func QueryStreetInfo(name string, pageNo int, pageSize int) string {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	street := table.Street{}
	street.Name = name
	finds := table.FindStreets(db, street, pageNo, pageSize)

	streets := table.Streets{finds}
	result, _ := json.Marshal(streets)
	return string(result)
}

//QueryComunityInfo 查询社区信息
func QueryComunityInfo(name string, pageNo int, pageSize int) string {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	community := table.Community{}
	community.Name = name

	finds := table.FindCommunities(db, community, pageNo, pageSize)

	if session != nil {
		session.Close()
	}

	communities := table.Communities{finds}
	result, _ := json.Marshal(communities)
	return string(result)
}

//QueryComunityDistinct 查询社区信息
func QueryComunityDistinct(key string) string {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)

	finds := table.FindCommunityDistincts(db, key)

	if session != nil {
		session.Close()
	}
	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryXiaoQuInfo 查询小区infos
func QueryXiaoQuInfo(name string, pageNo int, pageSize int) string {
	session, err := mgo.Dial(HOST)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	xiaoQu := table.XiaoQu{}
	xiaoQu.Name = name

	finds := table.FindXiaoQus(db, xiaoQu, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryXQDistinct 查询社区信息
func QueryXQDistinct(key string) string {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)

	finds := table.FindXQDistincts(db, key)

	if session != nil {
		session.Close()
	}
	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryGovInfo 查询小区infos
func QueryGovInfo(name string, pageNo int, pageSize int) string {
	session, err := mgo.Dial(HOST)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	gov := table.Gov{}
	gov.UserName = name

	finds := table.FindGovs(db, gov, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryStreetUserInfo 查询街道用户结果集
func QueryStreetUserInfo(name string, pageNo int, pageSize int) string {
	session, err := mgo.Dial(HOST)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	info := table.StreetUser{}
	info.UserName = name

	finds := table.FindStreetUsers(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryWXUser 查询街道用户结果集
func QueryWXUser(openID string, pageNo int, pageSize int) string {
	session, err := mgo.Dial(HOST)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	info := table.WXUser{}
	info.OpenID = openID

	finds := table.FindWXUsers(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryWuYe 查询街道用户结果集
func QueryWuYe(xiaoQu string, pageNo int, pageSize int) string {
	session, err := mgo.Dial(HOST)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	info := table.WuYe{}
	info.XiaoQu = xiaoQu

	finds := table.FindWuYes(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}

//QueryHouse 查询街道用户结果集
func QueryHouse(buildNo string, pageNo int, pageSize int) string {
	session, err := mgo.Dial(HOST)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBName)
	info := table.House{}
	info.BuildNo = buildNo

	finds := table.FindHouses(db, info, pageNo, pageSize)

	result, _ := json.Marshal(finds)
	return string(result)
}
