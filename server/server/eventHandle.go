package server

import (
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startEventHandle(router *gin.RouterGroup, dbc *mgo.Database) {
	// 按index 查询事件，如果index==“”则查询所有的
	router.POST("/eventHandles", func(c *gin.Context) {
		var queryInfo QueryEvent
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindEventHandle(dbc, queryInfo.Index,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/eventHandle/add", func(c *gin.Context) {
		var info table.EventHandle
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.InsertEventHandle(dbc, info))
		}
	})
	router.POST("/eventHandle/detail/:index", func(c *gin.Context) {
		index := c.Param("index")
		c.JSON(http.StatusOK, table.FindEventHandle(dbc, index,
			1, 1))
	})

	router.POST("/eventHandle/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventHandlesByKV(dbc, params))
		}
	})

	router.GET("/eventHandle/today/handles", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.FindTodayHandles(dbc))
	})
}
