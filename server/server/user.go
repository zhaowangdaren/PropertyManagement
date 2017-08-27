package server

import (
	"fmt"
	"net/http"
	"strconv"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startUser(router *gin.RouterGroup, dbc *mgo.Database) {
	router.GET("/users/num/:type", func(c *gin.Context) {
		userType, err := strconv.Atoi(c.Param("type"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, table.CountUsers(dbc, userType))
	})

	router.POST("/users", func(c *gin.Context) {
		var queryInfo QueryUser
		err := c.BindJSON(&queryInfo)
		fmt.Println("startUser", queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindUsers(dbc, queryInfo.UserName, queryInfo.Type,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/user/add", func(c *gin.Context) {
		var user table.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.InsertUser(dbc, user))
		}
	})

	//发送{values: [id,id,...]}，可支持删除多个id的用户
	router.POST("/users/del", func(c *gin.Context) {
		var ids Values
		err := c.BindJSON(&ids)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelUsers(dbc, ids.Values))
		}
	})

	router.POST("/user/update", func(c *gin.Context) {
		var user table.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.UpdateUser(dbc, user))
		}
	})

	router.POST("/user/changePassword", func(c *gin.Context) {
		var info ChangePassword
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.ChangePassword(dbc, info.Type, info.UserName, info.Password, info.NewPassword))
		}
	})
}
