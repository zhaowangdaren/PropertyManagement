package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"

	"../db/table"
)

func startCourtWX(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/court/wx", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindCourtByActualName(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/court/wx/update", func(c *gin.Context) {
		var info table.CourtWX
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.UpdateCourtWX(dbc, info))
		}
	})

	router.GET("/court/wx/del/openid/:openid", func(c *gin.Context) {
		openid := c.Param("openid")
		if openid == "" {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "openid为空"})
			return
		}
		err := table.DelCourtWX(dbc, openid)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
		}
	})
}

func startOpenCourt(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/court/wx/bind", func(c *gin.Context) {
		var info table.CourtWX
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1})
			glog.Error(err.Error())
		} else {
			info.Bind = 0
			c.JSON(http.StatusOK, table.InsertCourtWX(dbc, info))
		}
	})
}
