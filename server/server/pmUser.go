package server

import (
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startPMUser(router *gin.RouterGroup, dbc *mgo.Database) {
	//查询绑定的用户，入参为name、pageNo、pageSize
	router.POST("/pmUser", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindPMUsers(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	//获取所有key对应的不同value集合
	router.POST("/pmUser/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, table.FindPMUserDistinct(dbc, key))
	})
}
