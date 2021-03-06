package server

import (
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startPark(router *gin.RouterGroup, dbc *mgo.Database) {

	router.POST("/park/update", func(c *gin.Context) { // 更新停车位信息
		var obj table.Park
		err := c.BindJSON(&obj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.UpdatePark(dbc, obj))
		}
	})

	router.GET("/park/query/xqid/:xqid", func(c *gin.Context) {
		xqid := c.Param("xqid")
		c.JSON(http.StatusOK, table.FindParkByXQID(dbc, xqid))
	})

	router.POST("/park/managers", func(c *gin.Context) {
		var query QueryParkManager
		err := c.BindJSON(&query)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, table.FindPrakManagerByActualName(dbc,
			query.ActualName, query.PageNo, query.PageSize))
	})

	router.POST("/park/managers/del", func(c *gin.Context) {
		var values Values
		err := c.BindJSON(&values)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, table.DelParkManager(dbc, values.Values))
	})

	router.POST("/park/manager/bind", func(c *gin.Context) {
		var obj table.ParkManager
		err := c.BindJSON(&obj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		if obj.OpenID != "" && obj.ActualName != "" && obj.Tel != "" {
			obj.Bind = 0
			c.JSON(http.StatusOK, table.InsertParkManager(dbc, obj))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "信息错误"})
		}
	})

	router.POST("/park/manager/update", func(c *gin.Context) {
		var obj table.ParkManager
		err := c.BindJSON(&obj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, table.UpdateParkManager(dbc, obj))
	})

	router.POST("/park/manager/query/openid", func(c *gin.Context) {
		query := make(map[string]string)
		err := c.BindJSON(&query)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		openid := query["openid"]
		c.JSON(http.StatusOK, table.FindParkManagerByOpenID(dbc, openid))
	})
}
