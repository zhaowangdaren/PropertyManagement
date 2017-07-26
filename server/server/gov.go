package server

import (
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startGov(router *gin.Engine, dbc *mgo.Database) {
	router.POST("/gov", func(c *gin.Context) {
		var queryInfo QueryGov
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindGovs(dbc, queryInfo.UserName,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})
}
