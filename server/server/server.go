package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
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
		var form LoginForm
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, db.Login(dbc, form.User, form.Password))
		}
	})

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

	//localhost:3000/communityInfo?name=社区名&pageNo=1&pageSize=1
	router.POST("/community", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, db.QueryComunityInfo(dbc, queryInfo.Name, queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		}
	})

	router.POST("/community/add", func(c *gin.Context) {
		var jsonComm table.Community
		err := c.BindJSON(&jsonComm)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, table.InsertCommunity(dbc, jsonComm))
		}
	})

	//按community的name值删数据
	router.POST("/community/del", func(c *gin.Context) {
		var names Values
		err := c.BindJSON(&names)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelCommunities(dbc, names.Values))
		}
	})
	//localhost:3000/communityInfo/key?key=name
	router.POST("/community/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, db.QueryComunityDistinct(dbc, key))
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

		// queryStr := c.PostForm("query")
		// var query map[string]interface{}
		// err := json.Unmarshal([]byte(queryStr), &query)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println("query", query)
		// result := db.QueryComunityKVs(dbc, query)
		// c.String(http.StatusOK, result)
	})

	router.POST("/community/streetName/:streetName", func(c *gin.Context) {
		streetName := c.Param("streetName")
		log.Println("streetName", streetName)
		c.JSON(http.StatusOK, table.FindCommunityByStreetName(dbc, streetName))
	})

	router.POST("/xq", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, db.QueryXiaoQuInfo(dbc, queryInfo.Name, queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		}
	})

	router.POST("/xq/add", func(c *gin.Context) {
		var jsonObj table.XiaoQu
		err := c.BindJSON(&jsonObj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "params error"})
			log.Print(err)
		} else {
			c.JSON(http.StatusOK, table.InsertXQ(dbc, jsonObj))
		}
	})

	//按xq的name值删数据，传入数据为json:{values: []}
	router.POST("/xq/del", func(c *gin.Context) {
		var names Values
		err := c.BindJSON(&names)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelXQs(dbc, names.Values))
		}
	})

	//localhost:3000/xiaoQu/key?key=name
	router.POST("/xq/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, db.QueryXQDistinct(dbc, key))
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

	//TODO 参照上面的xq的查、增加、删除，为下面的API增加相应接口
	startGov(router, dbc)

	router.POST("/streetUser", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindStreetUsers(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/streetUser/add", func(c *gin.Context) {
		var jsonObj table.StreetUser
		err := c.BindJSON(&jsonObj)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Print(err)
		} else {
			c.JSON(http.StatusOK, table.InsertStreetUser(dbc, jsonObj))
		}
	})

	//查询绑定的用户，入参为name、pageNo、pageSize
	router.POST("/wxUser", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindWXUsers(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	//获取所有key对应的不同value集合
	router.POST("/wxUser/key/:key", func(c *gin.Context) {
		key := c.Param("key")
		c.JSON(http.StatusOK, table.FindWXUserDistinct(dbc, key))
	})

	router.POST("/pm", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindWuYes(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/pm/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindWuYeKVs(dbc, params))
		}
	})

	router.POST("/house", func(c *gin.Context) {
		var queryInfo QueryHouse
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindHouses(dbc, queryInfo.BuildNo,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})
	startLogin(router, dbc)
	startEvent(router, dbc)
	startEventHandle(router, dbc)
	startPMKPI(router, dbc)
	router.Run(":3000")
}
