package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"

	"../db/table"
)

func startAnnounce(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/announce/search/name", func(c *gin.Context) {
		var query QueryBasic
		err := c.BindJSON(&query)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		fmt.Println("announce", query.PageSize)
		result, err := table.FindAnnounce(dbc, query.Name, query.PageNo, query.PageSize)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
	})
	router.POST("/announce/add", func(c *gin.Context) {
		var announces AddAnnounces
		err := c.BindJSON(&announces)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.InsertAnnounces(dbc, announces.List)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
	})

	router.GET("/announce/del/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := table.DelAnnounce(dbc, id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
	})
}
