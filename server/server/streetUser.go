package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startStreetUser(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/streetUser", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindStreetUsers(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/streetUser/add", func(c *gin.Context) {
		var jsonObj table.StreetUser
		err := c.BindJSON(&jsonObj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Print(err)
		} else {
			c.JSON(http.StatusOK, table.InsertStreetUser(dbc, jsonObj))
		}
	})
}
