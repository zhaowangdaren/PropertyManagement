package server

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"../db/table"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

func PostJson(url, post string) {
	fmt.Println("PostJson", post)
	var jsonStr = []byte(post)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		glog.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Error(err.Error())
	}
	fmt.Println("PostJson", resp)
	defer resp.Body.Close()
}

func FilterEventStatus(status int) string {
	switch status {
	case -1:
		return "用户撤销"
	case 0:
		return "居民提交"
	case 1:
		return "已审核待处理"
	case 2:
		return "已解决"
	case 3:
		return "已关闭"
	default:
		return ""
	}
}

func FilterEventHandleAuthorCategory(authorCategory int) string {
	switch authorCategory {
	case 1:
		return "系统管理员"
	case 2:
		return "政府主管单位"
	case 3:
		return "街道办公人员"
	}
	return ""
}
func PushNotice2WX(eventHandle table.EventHandle, dbc *mgo.Database) {

	event := table.FindEvent(dbc, eventHandle.Index)
	userOpenID := event.OpenID
	if userOpenID == "" {
		glog.Error("Event:" + eventHandle.Index + " 没有查询到其OpenID")
		return
	}
	xqName := table.FindXQ(dbc, event.XQID).Name
	time := time.Unix(event.Time, 0).Format("2006-01-02 15:04:05")
	post := "{" +
		"\"touser\": \"" + userOpenID + "\"," +
		"\"template_id\": \"JfYCUICcZxvOjdYYFUQVVu47AepqfhGau0nvLhGPcVA\"," +
		"\"url\": \"https://www.maszfglzx.com/#/wx/detailsProgress?index=" + eventHandle.Index + "&status=" + string(event.Status) + "\"," +
		"\"data\": {" +
		"\"first\": {" +
		"\"value\": \"您好，您的投诉已处理。具体信息如下：\"}," +
		"\"keyword1\": {" +
		"\"value\": \"" + xqName + "\"}," +
		"\"keyword2\": {" +
		"\"value\": \"" + FilterEventStatus(event.Status) + "\"}," +
		"\"keyword3\": {" +
		"\"value\": \"" + time + "\"}," +
		"\"remark\": {" +
		"\"value\": \"感谢您的反馈\"}" +
		"}}"
	access_token := GetAccessToken()
	wxURL := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token
	PostJson(wxURL, post)
}

func startEventHandle(router *gin.RouterGroup, dbc *mgo.Database) {
	// 按index 查询事件，如果index==“”则查询所有的
	router.POST("/eventHandles", func(c *gin.Context) {
		var queryInfo QueryEvent
		err := c.BindJSON(&queryInfo)
		if err == nil {
			result := table.FindEventHandle(dbc, queryInfo.Index,
				queryInfo.PageNo, queryInfo.PageSize)
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/eventHandle/add", func(c *gin.Context) {
		var info table.EventHandle
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			result, err := table.InsertEventHandle(dbc, info)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
			PushNotice2WX(info, dbc)
		}
	})
	router.POST("/eventHandle/detail/:index", func(c *gin.Context) {
		index := c.Param("index")
		c.JSON(http.StatusOK, table.FindEventHandle(dbc, index,
			1, 1))
	})

	router.POST("/eventHandle/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventHandlesByKV(dbc, params))
		}
	})

	router.GET("/eventHandle/today/handles", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.FindTodayHandles(dbc))
	})
}
