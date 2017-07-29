package server

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"time"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

const TenMinute = 10 * 60 * 1000

func token(c *gin.Context, userName string, password string, userType int) table.OnlineUser {
	clientIP := c.ClientIP()
	timeNow := time.Now().Unix()
	md5Token := md5.New()
	lastStr := userName + "-" + clientIP + "-" + string(timeNow)
	io.WriteString(md5Token, lastStr)
	userToken := fmt.Sprintf("%x", md5Token.Sum([]byte(password)))
	result := table.OnlineUser{userName, userType, timeNow, userToken, TenMinute}
	return result
}

func startLogin(router *gin.Engine, dbc *mgo.Database) {
	router.POST("/admin/login", func(c *gin.Context) {
		var user table.Admin
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		if user.UserName == "" {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "用户名为空"})
			return
		}
		if user.Password == "" {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "密码为空"})
			return
		}
		result := table.FindAdmin(dbc, user.UserName)
		if result.UserName == user.UserName && result.Password == user.Password {
			onlineUser := token(c, user.UserName, user.Password, 0)
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": onlineUser.Token})
			table.UpsertOnlineUser(dbc, onlineUser)
			return
		} else if result.UserName != user.UserName {
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": "用户名不存在"})
			return
		} else if result.Password != user.Password {
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": "密码错误"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": "登录失败"})
	})

	router.POST("/street/login", func(c *gin.Context) {

	})

	router.POST("/gov/login", func(c *gin.Context) {

	})
}
