package server

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	// "github.com/gin-gonic/autotls"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	mgo "gopkg.in/mgo.v2"

	"../db"
	"../db/table"
	"../jwt"
)

func mapParams(params url.Values) map[string]string {
	result := make(map[string]string)
	for k, v := range params {
		result[string(k)] = strings.Join(v, "")
	}
	return result
}

//Start 使用Gin Web Framework
func Start() {
	flag.Parse()
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	dbc := session.DB(db.DBName)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins: []string{"http://localhost:9000", "https://localhost:9000", "http://10.176.115.42:9000", "http://10.176.118.61:9000", "https://127.0.0.1:9000", "103.238.215.206"},
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// router.Use(cors.Default())
	authMid := &jwt.GinJWTMiddleware{
		// Realm:      ".localhost:9000",
		Realm:      "test zone",
		Key:        []byte("123456"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userName string, password string, userType int, c *gin.Context) (table.User, bool) {
			result := table.FindUser(dbc, userName, userType)
			fmt.Println("Authenticator", result)
			if userName == result.UserName && password == result.Password && userType == result.Type && result.Pass == 1 {
				return result, true
			}
			return result, false
		},
		Authorizator: func(userID string, userType int, c *gin.Context) bool {
			// fmt.Println("userType", userType)
			// fmt.Println("Path", c.Request.URL.Path)
			return AuthControl(c.Request.URL.Path, userType)
			// return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"error": 1,
				"data": gin.H{
					"code":    code,
					"message": message,
				},
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	router.POST("/login", authMid.LoginHandler)
	authorized := router.Group("/pm")

	authorized.Use(authMid.MiddlewareFunc())
	{
		authorized.POST("/refresh_token", authMid.RefreshHandler)
		startUser(authorized, dbc)
		startStreet(authorized, dbc)
		startCommunity(authorized, dbc)
		startXQ(authorized, dbc)
		startPMUser(authorized, dbc)
		startPM(authorized, dbc)
		startHouse(authorized, dbc)
		startEvent(authorized, dbc)
		startEventHandle(authorized, dbc)
		startPMKPI(authorized, dbc)
		startPark(authorized, dbc)
	}

	open := router.Group("/open")
	{
		startOpen(open, dbc)
	}
	router.Run(":3000")
	// router.RunTLS(":3000", "/etc/letsencrypt/live/maszfglzx.com/fullchain.pem", "/etc/letsencrypt/live/maszfglzx.com/privkey.pem")
	glog.Flush()
}
