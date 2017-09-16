package server

import (
	"net/http"
	"strconv"

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

	router.GET("/event/del/index/:index", func(c *gin.Context) {
		index := c.Param("index")
		err := table.DelEvent(dbc, index)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
		}
	})

	router.POST("/event/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventKVs(dbc, params))
		}
	})

	router.POST("/event/kvs/page", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventKVsPage(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})

	router.POST("/event/kvs/page/gov", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventKVsPageForGov(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})

	// 查询当个kv，并分页显示
	router.POST("/event/kv", func(c *gin.Context) {
		var params QueryKV
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK,
				table.FindEventKV(dbc,
					params.Key, params.Value, params.PageNo, params.PageSize))
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
	//删除事件类型
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
		maxRecs, err := strconv.Atoi(c.Query("recs"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, table.CountDiffKeyEvents(dbc, key, maxRecs))
	})

	router.GET("/event/today/events", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.FindTodayEvents(dbc))
	})

	router.GET("/events/num", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.CountEvents(dbc))
	})
}
