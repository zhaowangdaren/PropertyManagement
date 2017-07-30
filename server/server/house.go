package server

import (
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startHouse(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/house", func(c *gin.Context) {
		var queryInfo QueryHouse
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindHouses(dbc, queryInfo.BuildNo,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})
}
