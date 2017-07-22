package server

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"../db"
	"github.com/gin-gonic/gin"
)

func mapParams(params url.Values) map[string]string {
	result := make(map[string]string)
	for k, v := range params {
		result[string(k)] = strings.Join(v, "")
	}
	return result
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

//Start 使用Gin Web Framework
func Start() {
	router := gin.Default()
	router.Use(cors())
	router.POST("/login", func(c *gin.Context) {
		userName := c.Query("userName")
		password := c.Query("password")
		c.String(http.StatusOK, db.Login(userName, password))
	})

	router.POST("/streetInfo", func(c *gin.Context) {
		name := c.Query("name")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryStreetInfo(name, pageNo, pageSize))
	})

	//localhost:3000/communityInfo?name=社区名&pageNo=1&pageSize=1
	router.POST("/communityInfo", func(c *gin.Context) {
		name := c.Query("name")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryComunityInfo(name, pageNo, pageSize))
	})

	//localhost:3000/communityInfo/key?key=name
	router.POST("/communityInfo/key", func(c *gin.Context) {
		key := c.Query("key")
		c.String(http.StatusOK, db.QueryComunityDistinct(key))
	})

	router.POST("/xiaoQu", func(c *gin.Context) {
		name := c.Query("name")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryXiaoQuInfo(name, pageNo, pageSize))
	})

	//localhost:3000/xiaoQu/key?key=name
	router.POST("/xiaoQu/key", func(c *gin.Context) {
		key := c.Query("key")
		c.String(http.StatusOK, db.QueryXQDistinct(key))
	})

	router.POST("/gov", func(c *gin.Context) {
		username := c.Query("username")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryGovInfo(username, pageNo, pageSize))
	})

	router.POST("/streetUser", func(c *gin.Context) {
		username := c.Query("username")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryStreetUserInfo(username, pageNo, pageSize))
	})

	router.POST("/wx", func(c *gin.Context) {
		openID := c.Query("openID")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryWXUser(openID, pageNo, pageSize))
	})

	router.POST("/wuye", func(c *gin.Context) {
		xiaoQu := c.Query("xiaoQu")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryWuYe(xiaoQu, pageNo, pageSize))
	})

	router.POST("/house", func(c *gin.Context) {
		buildNo := c.Query("buildNo")
		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		c.String(http.StatusOK, db.QueryHouse(buildNo, pageNo, pageSize))
	})
	router.Run(":3000")
}
