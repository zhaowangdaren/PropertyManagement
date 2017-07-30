package server

import (
	"fmt"
	"log"
	"net/http"

	"../db"
	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startStreet(router *gin.RouterGroup, dbc *mgo.Database) {
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
		// name := c.PostForm("name")
		// pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		// pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		// c.String(http.StatusOK, db.QueryStreetInfo(dbc, name, pageNo, pageSize))
	})

	router.POST("/street/add", func(c *gin.Context) {
		var jsonStreet table.Street
		err := c.BindJSON(&jsonStreet)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "params error"})
			panic(err)
		} else {
			c.JSON(http.StatusOK, table.InsertStreet(dbc, jsonStreet))
		}
	})
	//按street的name删数据
	router.POST("/street/del", func(c *gin.Context) {
		var names Values
		err := c.BindJSON(&names)
		fmt.Println("/street/del names", names.Values)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelStreets(dbc, names.Values))
		}
	})
	//筛选出key-value中value不重复的
	router.POST("/street/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, table.FindStreetDistinct(dbc, key))
	})
}
