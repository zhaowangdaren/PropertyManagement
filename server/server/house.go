package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
)

func startHouse(router *gin.RouterGroup, dbc *mgo.Database) {
	router.GET("/houses/num", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.CountHouses(dbc))
	})

	router.POST("/house", func(c *gin.Context) {
		var queryInfo QueryHouse
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindHouses(dbc, queryInfo.BuildNo,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/house/add", func(c *gin.Context) {
		var add table.House
		err := c.BindJSON(&add)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			panic(err)
		} else {
			c.JSON(http.StatusOK, table.InsertHouse(dbc, add))
		}
	})
	router.POST("/house/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindHouseKVs(dbc, params))
		}
	})

	router.POST("/house/kvs/page", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindHouseByKVsPage(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})
	router.POST("/house/update", func(c *gin.Context) {
		var update table.House
		err := c.BindJSON(&update)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "params error"})
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, table.UpdateHouse(dbc, update))
		}
	})

	//按street的name删数据, 删除多个
	router.POST("/houses/del", func(c *gin.Context) {
		var ids Values
		err := c.BindJSON(&ids)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelHouses(dbc, ids.Values))
		}
	})

	// 导入xml，从xml中读取建筑信息
	router.POST("/house/import/xml", func(c *gin.Context) {
		var files Values
		err := c.BindJSON(&files)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		importInfo := ""
		for _, fileName := range files.Values {
			err = table.ImportHousesFromXML(dbc, FileBasicPath+"/files/"+fileName)
			// if err != nil {
			// 	c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			// } else {
			// }
			if err != nil {
				glog.Error(err.Error())
				importInfo += err.Error()
			}
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": "导入完成," + importInfo})
	})
}
