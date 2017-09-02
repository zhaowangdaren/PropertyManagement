package server

import (
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
)

func startOpenParkManager(router *gin.RouterGroup, dbc *mgo.Database) {
	// 新增停车管理员
	router.POST("/park/manager/insert", func(c *gin.Context) {
		glog.Info(c)
		var obj table.ParkManager
		err := c.BindJSON(&obj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.InsertParkManager(dbc, obj))
		}
	})
}
