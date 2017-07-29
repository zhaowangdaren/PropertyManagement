package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startPMKPI(router *gin.Engine, dbc *mgo.Database) {
	router.POST("/pmkpi", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindPMKPIsByName(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/pmkpi/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindPMKPIByKVs(dbc, params))
		}
	})
}
