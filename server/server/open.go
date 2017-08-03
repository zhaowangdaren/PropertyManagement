//Package Server 公开API，例如供微信端用户投诉，查询街道
package server

import (
	"log"
	"net/http"

	"../db"
	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startOpen(router *gin.RouterGroup, dbc *mgo.Database) {
	//查询街道信息，入参：name\pageNo\pageSize，当name为空时会查询所有的街道
	router.POST("/street", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, db.QueryStreetInfo(dbc, queryInfo.Name, queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		}
	})

	router.POST("/community/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindCommunitiesKV(dbc, params))
		}
	})

	router.POST("/xq/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindXQKVs(dbc, params))
		}
	})
}
