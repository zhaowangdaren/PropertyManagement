package server

import (
	"log"
	"net/http"

	"../db"
	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startXQ(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/xq", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, db.QueryXiaoQuInfo(dbc, queryInfo.Name, queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		}
	})

	router.POST("/xq/add", func(c *gin.Context) {
		var jsonObj table.XiaoQu
		err := c.BindJSON(&jsonObj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "params error"})
			log.Print(err)
		} else {
			c.JSON(http.StatusOK, table.InsertXQ(dbc, jsonObj))
		}
	})

	router.POST("/xq/update", func(c *gin.Context) {
		var jsonObj table.XiaoQu
		err := c.BindJSON(&jsonObj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "params error"})
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, table.UpdateXQ(dbc, jsonObj))
		}
	})

	//按xq的name值删数据，传入数据为json:{values: []}
	router.POST("/xq/del", func(c *gin.Context) {
		var names Values
		err := c.BindJSON(&names)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelXQs(dbc, names.Values))
		}
	})

	//localhost:3000/xiaoQu/key?key=name
	router.POST("/xq/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, db.QueryXQDistinct(dbc, key))
	})

	router.POST("/xq/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindXQKVs(dbc, params))
		}
	})

	//通过ID数组查找街道，返回的信息也是数组
	router.POST("/xq/ids", func(c *gin.Context) {
		var ids Values
		err := c.BindJSON(&ids)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindXQsByIDs(dbc, ids.Values))
		}
	})
}
