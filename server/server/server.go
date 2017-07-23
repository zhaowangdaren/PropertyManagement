package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"

	"../db"
	"../db/table"
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
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	dbc := session.DB(db.DBName)

	router := gin.Default()
	router.Use(cors())
	router.POST("/login", func(c *gin.Context) {
		userName := c.PostForm("userName")
		password := c.PostForm("password")
		c.String(http.StatusOK, db.Login(dbc, userName, password))
	})

	router.POST("/streetInfo", func(c *gin.Context) {
		var queryInfo QueryStreet
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": db.QueryStreetInfo(dbc, queryInfo.Name, queryInfo.PageNO, queryInfo.PageSize)})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "params error"})
			panic(err)
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
			c.JSON(http.StatusOK, db.InsertStreet(dbc, jsonStreet))
		}
	})

	//localhost:3000/communityInfo?name=社区名&pageNo=1&pageSize=1
	router.POST("/communityInfo", func(c *gin.Context) {
		name := c.PostForm("name")
		pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		fmt.Println("name", name)
		c.String(http.StatusOK, db.QueryComunityInfo(dbc, name, pageNo, pageSize))
	})

	//localhost:3000/communityInfo/key?key=name
	router.POST("/communityInfo/key", func(c *gin.Context) {
		key := c.PostForm("key")
		c.String(http.StatusOK, db.QueryComunityDistinct(dbc, key))
	})

	router.POST("/communityKVs", func(c *gin.Context) {
		queryStr := c.PostForm("query")
		var query map[string]interface{}
		err := json.Unmarshal([]byte(queryStr), &query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("query", query)
		result := db.QueryComunityKVs(dbc, query)
		c.String(http.StatusOK, result)
	})
	router.POST("/xiaoQu", func(c *gin.Context) {
		name := c.PostForm("name")
		pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		c.String(http.StatusOK, db.QueryXiaoQuInfo(dbc, name, pageNo, pageSize))
	})

	//localhost:3000/xiaoQu/key?key=name
	router.POST("/xiaoQu/key", func(c *gin.Context) {
		key := c.PostForm("key")
		c.String(http.StatusOK, db.QueryXQDistinct(dbc, key))
	})

	router.POST("/xiaoQuKVs", func(c *gin.Context) {
		queryStr := c.PostForm("query")
		var query map[string]interface{}
		err := json.Unmarshal([]byte(queryStr), &query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("query", query)
		result := db.QueryXQKVs(dbc, query)
		c.String(http.StatusOK, result)
	})

	router.POST("/gov", func(c *gin.Context) {
		username := c.PostForm("username")
		pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		c.String(http.StatusOK, db.QueryGovInfo(dbc, username, pageNo, pageSize))
	})

	router.POST("/streetUser", func(c *gin.Context) {
		username := c.PostForm("username")
		pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		c.String(http.StatusOK, db.QueryStreetUserInfo(dbc, username, pageNo, pageSize))
	})

	router.POST("/wx", func(c *gin.Context) {
		openID := c.PostForm("openID")
		pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		c.String(http.StatusOK, db.QueryWXUser(dbc, openID, pageNo, pageSize))
	})

	router.POST("/wuye", func(c *gin.Context) {
		xiaoQu := c.PostForm("xiaoQu")
		pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		c.String(http.StatusOK, db.QueryWuYe(dbc, xiaoQu, pageNo, pageSize))
	})

	router.POST("/house", func(c *gin.Context) {
		buildNo := c.PostForm("buildNo")
		pageNo, _ := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
		c.String(http.StatusOK, db.QueryHouse(dbc, buildNo, pageNo, pageSize))
	})
	router.Run(":3000")
}
