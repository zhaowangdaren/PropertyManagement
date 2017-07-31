package server

import (
	"log"
	"net/http"

	"../db"
	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startCommunity(router *gin.RouterGroup, dbc *mgo.Database) {
	//localhost:3000/communityInfo?name=社区名&pageNo=1&pageSize=1
	router.POST("/community", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, db.QueryComunityInfo(dbc, queryInfo.Name, queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		}
	})

	router.POST("/community/add", func(c *gin.Context) {
		var jsonComm table.Community
		err := c.BindJSON(&jsonComm)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, table.InsertCommunity(dbc, jsonComm))
		}
	})

	//按community的name值删数据
	router.POST("/community/del", func(c *gin.Context) {
		var names Values
		err := c.BindJSON(&names)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelCommunities(dbc, names.Values))
		}
	})

	router.POST("/community/update", func(c *gin.Context) {
		var jsonObj table.Community
		err := c.BindJSON(&jsonObj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "params error"})
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, table.UpdateCommunity(dbc, jsonObj))
		}
	})
	//localhost:3000/communityInfo/key?key=name
	router.POST("/community/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, db.QueryComunityDistinct(dbc, key))
	})

	router.POST("/community/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindCommunitiesKV(dbc, params))
		}
	})

	router.POST("/community/streetName/:streetName", func(c *gin.Context) {
		streetName := c.Param("streetName")
		log.Println("streetName", streetName)
		c.JSON(http.StatusOK, table.FindCommunityByStreetName(dbc, streetName))
	})
}
