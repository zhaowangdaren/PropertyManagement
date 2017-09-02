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
}
