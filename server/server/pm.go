package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startPM(router *gin.RouterGroup, dbc *mgo.Database) {
	router.GET("/pms/num", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.CountPMs(dbc))
	})
	router.POST("/pm", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindWuYes(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/pm/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindWuYeKVs(dbc, params))
		}
	})

	router.POST("/pm/add", func(c *gin.Context) {
		var pm table.PM
		err := c.BindJSON(&pm)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.InsertPM(dbc, pm))
		}
	})

	router.POST("/pm/update", func(c *gin.Context) {
		var update table.PM
		err := c.BindJSON(&update)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, table.UpdatePM(dbc, update))
		}
	})

	//按street的name删数据, 删除多个
	router.POST("/pms/del", func(c *gin.Context) {
		var names Values
		err := c.BindJSON(&names)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelPMs(dbc, names.Values))
		}
	})
}
