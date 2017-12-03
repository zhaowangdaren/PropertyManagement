package server

import (
	"fmt"
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func startAnnounce(router *gin.RouterGroup, dbc *mgo.Database) {
	router.POST("/announce/search/name", func(c *gin.Context) {
		var query QueryBasic
		err := c.BindJSON(&query)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		fmt.Println("announce", query.PageSize)
		result, err := table.FindAnnounce(dbc, query.Name, query.PageNo, query.PageSize)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
	})
	router.POST("/announce/add", func(c *gin.Context) {
		var announces AddAnnounces
		err := c.BindJSON(&announces)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.InsertAnnounces(dbc, announces.List)
		if err.Error() != "" {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
		for _, announce := range announces.List {
			noticeAllPM(dbc, announce.FileName)
		}
	})

	router.GET("/announce/del/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := table.DelAnnounce(dbc, id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
	})
}

func noticeAllPM(dbc *mgo.Database, fileName string) {
	kvs := make(map[string]interface{})
	kvs["bind"] = 1
	pmUsers := table.FindPMUserByKVs(dbc, kvs)
	var announcedPM []string
	for _, pmUser := range pmUsers {
		hasSend := Contains(announcedPM, pmUser.OpenID)
		if hasSend {
			continue
		}
		fmt.Println("OpenID", pmUser.OpenID)
		announcedPM = append(announcedPM, pmUser.OpenID)
		pjson := `{
			"touser": "` + pmUser.OpenID + `",
			"template_id": "_yYzGHjaSbNh9FqugaRQjE-l_56szVL05X3tM-XDgds",
			"url": "",
			"data": {
				"first": {
					"value": "您好，政府部门发布了新的公告：` + fileName + `"
				},
				"keyword1": {
					"value": ""
				},
				"keyword2": {
					"value": "` + fileName + `"
				},
				"keyword3": {
					"value": "政府部门"
				},
				"keyword4": {
					"value": ""
				},
				"remark": {
					"value": "请及时登陆官网查看新公告"
				}
			}
		}`
		accessToken := GetAccessToken()
		wxURL := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + accessToken
		PostJson(wxURL, pjson)
	}
}

// Contain 判断obj是否在target中，target支持的类型arrary,slice,map
func Contains(slice []string, item string) bool {
	if len(slice) == 0 {
		return false
	}
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
