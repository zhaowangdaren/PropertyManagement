package server

import (
	"log"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startPMKPI(router *gin.RouterGroup, dbc *mgo.Database) {
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

	router.POST("/pmkpi/kvs/page", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindPMKPIByKVsPage(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})
}
