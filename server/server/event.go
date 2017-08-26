package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startEvent(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/event", func(c *gin.Context) {
		var queryInfo QueryEvent
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindEvents(dbc, queryInfo.Index,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/event/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindEventKVs(dbc, params))
		}
	})

	router.POST("/event/update", func(c *gin.Context) {
		var info table.Event
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.UpdateEvent(dbc, info))
		}
	})

	router.POST("/event/type/add", func(c *gin.Context) {
		var info table.EventType
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.InsertEventType(dbc, info.Type))
		}
	})

	router.POST("/event/type/del", func(c *gin.Context) {
		var info table.EventType
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelEventType(dbc, info.Type))
		}
	})

	router.GET("/event/key/nums/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, table.CountDiffKeyEvents(dbc, key))
	})
}
