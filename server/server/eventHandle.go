package server

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	defer resp.Body.Close()
}

func FilterEventStatus(status int) string {
	switch status {
	case -2:
		return "已关闭"
	case -1:
		return "已撤销"
	case 0:
		return "居民提交"
	case 1:
		return "处理中"
	case 2:
		return "已处理待确认"
	case 3:
		return "已解决"
	case 4:
		return "未解决"
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
	if eventHandle.HandleType > 0 && event.Status == 0 {
		event.Status = 1
		table.UpdateEventStatus(dbc, event.Index, event.Status)
	}
	userOpenID := event.OpenID
	if userOpenID == "" {
		glog.Error("Event:" + eventHandle.Index + " 没有查询到其OpenID")
		return
	}
	xqName := table.FindXQ(dbc, event.XQID).Name
	time := time.Unix(event.Time, 0).Format("2006-01-02 15:04:05")
	status := FilterEventStatus(event.Status)
	// good
	// 	pjson := `{
	//   "touser": "oIVRq0ubSS9zMeCAKE55hlIAFBj8",
	//   "template_id": "JfYCUICcZxvOjdYYFUQVVu47AepqfhGau0nvLhGPcVA",
	//   "url": "https://www.maszfglzx.com/#/wx/detailsProgress?index=20170907221818572142744&status=",
	//   "data": {
	//     "first": {
	//       "value": "您好，您的投诉已处理。具体信息如下："
	//     },
	//     "keyword1": {
	//       "value": ""
	//     },
	//     "keyword2": {
	//       "value": "居民提交"
	//     },
	//     "keyword3": {
	//       "value": "2017-09-07 22:18:18"
	//     },
	//     "remark": {
	//       "value": "感谢您的反馈"
	//     }
	//   }
	// }`
	// https://www.maszfglzx.com/#/wx/detailsProgress?index=` + eventHandle.Index + `
	pjson := `{
  "touser": "` + userOpenID + `",
  "template_id": "JfYCUICcZxvOjdYYFUQVVu47AepqfhGau0nvLhGPcVA",
  "url": "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxa768bfacbb694944&redirect_uri=https%3A%2F%2Fwww.maszfglzx.com%2Fcomplaint-progress.html&response_type=code&scope=snsapi_base&state=` + eventHandle.Index + `#wechat_redirect",
  "data": {
    "first": {
      "value": "您好，您的投诉已处理。具体信息如下："
    },
    "keyword1": {
      "value": "` + xqName + `"
    },
		"keyword2": {
      "value": "` + time + `"
    },
    "keyword3": {
      "value": "` + status + `"
    },
		"keyword4": {
      "value": ""
    },
    "remark": {
      "value": "感谢您的反馈"
    }
  }
}`
	access_token := GetAccessToken()
	wxURL := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token
	PostJson(wxURL, pjson)
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

	router.POST("/eventHandle/kvs/page", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventHandleByKVsPage(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})

	router.GET("/eventHandle/today/handles", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.FindTodayHandles(dbc))
	})

	router.POST("/eventHandle/audit/level", func(c *gin.Context) {
		var eventHandle table.EventHandle
		err := c.BindJSON(&eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventHandle.HandleType = 7
		_, err = table.InsertEventHandle(dbc, eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.UpdateEventLevel(dbc, eventHandle.Index, eventHandle.EventLevel)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
	})

	router.POST("/eventHandle/request/close", func(c *gin.Context) {
		var eventHandle table.EventHandle
		err := c.BindJSON(&eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventHandle.HandleType = 8
		_, err = table.InsertEventHandle(dbc, eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.UpdateEventRequestClose(dbc, eventHandle.Index, 1)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
	})

	router.POST("/eventHandle/agree/close", func(c *gin.Context) {
		var eventHandle table.EventHandle
		err := c.BindJSON(&eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventHandle.HandleType = 9
		_, err = table.InsertEventHandle(dbc, eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.UpdateEventStatus(dbc, eventHandle.Index, -2)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
	})

	router.POST("/eventHandle/notice/court", func(c *gin.Context) {
		var eventHandle table.EventHandle
		err := c.BindJSON(&eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventHandle.HandleType = 11
		_, err = table.InsertEventHandle(dbc, eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.UpdateEventToCourt(dbc, eventHandle.Index, 1)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})

		// push to WeChat
		xqName := table.FindXQ(dbc, eventHandle.XQID).Name
		kvs := make(map[string]interface{})
		kvs["bind"] = 1
		courts := table.FindCourtByKVs(dbc, kvs)
		PushNotice2Court(courts, xqName, eventHandle.Index)
	})

	router.POST("/eventHandle/notice/gov", func(c *gin.Context) {
		var eventHandle table.EventHandle
		err := c.BindJSON(&eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventHandle.HandleType = 10
		_, err = table.InsertEventHandle(dbc, eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.UpdateEventNoticeGov(dbc, eventHandle.Index, 1)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
	})

	router.POST("/eventHandle/noticepm", func(c *gin.Context) {
		var eventHandle table.EventHandle
		err := c.BindJSON(&eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventHandle.HandleType = 1
		eventIndex := eventHandle.Index
		xqid := eventHandle.XQID
		pm := table.FindPMByKV(dbc, "xqid", xqid)
		// fmt.Println("pmID:", pm.ID.Hex())
		// pmUser := table.FindPMUserByKV(dbc, "pmid", pm.ID.Hex())
		kvs := make(map[string]interface{})
		kvs["pmid"] = pm.ID.Hex()
		kvs["bind"] = 1
		pmUsers := table.FindPMUserByKVs(dbc, kvs)
		if len(pmUsers) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "该小区物业人员未绑定微信"})
			return
		}
		event := table.FindEvent(dbc, eventIndex)
		xqName := table.FindXQ(dbc, event.XQID).Name
		PushNotice2PM(pmUsers, xqName, eventIndex, pm.Name)
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})

		table.InsertEventHandle(dbc, eventHandle)
	})

	router.POST("/eventHandle/govtalkpm", func(c *gin.Context) {
		var eventHandle table.EventHandle
		err := c.BindJSON(&eventHandle)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventHandle.HandleType = 2
		pm := table.FindPMByKV(dbc, "xqid", eventHandle.XQID)
		kvs := make(map[string]interface{})
		kvs["pmid"] = pm.ID.Hex()
		kvs["bind"] = 1
		pmUsers := table.FindPMUserByKVs(dbc, kvs)
		if len(pmUsers) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "该小区物业人员未绑定微信"})
			return
		}
		xqName := table.FindXQ(dbc, eventHandle.XQID).Name
		GOVTalkPM(pmUsers, xqName, eventHandle.Index, pm.Name, eventHandle.HandleInfo)
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})

		table.InsertEventHandle(dbc, eventHandle)
		table.UpdateEventTalkAbout(dbc, eventHandle.Index, 1)
	})
}

//NoticePM 通知PM
func NoticePM(dbc *mgo.Database, eventIndex string) error {
	event := table.FindEvent(dbc, eventIndex)
	pm := table.FindPMByKV(dbc, "xqid", event.XQID)
	// fmt.Println("pmID:", pm.ID.Hex())
	// pmUser := table.FindPMUserByKV(dbc, "pmid", pm.ID.Hex())
	kvs := make(map[string]interface{})
	kvs["pmid"] = pm.ID.Hex()
	kvs["bind"] = 1
	pmUsers := table.FindPMUserByKVs(dbc, kvs)
	if len(pmUsers) == 0 {
		return errors.New("该小区物业人员未绑定微信")
	}

	xqName := table.FindXQ(dbc, event.XQID).Name
	PushNotice2PM(pmUsers, xqName, eventIndex, pm.Name)
	return nil
}

func startOpenEventHandle(router *gin.RouterGroup, dbc *mgo.Database) {
	// router.POST("/eventHandle/kvs", func(c *gin.Context) {
	// 	params := make(map[string]interface{})
	// 	err := c.BindJSON(&params)
	// 	if err != nil {
	// 		c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
	// 	} else {
	// 		c.JSON(http.StatusOK, table.FindEventHandlesByKV(dbc, params))
	// 	}
	// })
	router.POST("/eventHandle/kvs/page", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventHandleByKVsPage(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})

	router.POST("/eventHandle/pm/deal", func(c *gin.Context) {
		var info table.EventHandle
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		info.HandleType = 5
		result, err := table.InsertEventHandle(dbc, info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
		table.UpdateEventStatus(dbc, info.Index, 2)
		PushNotice2WX(info, dbc)
	})

	router.POST("/eventHandle/user/reply", func(c *gin.Context) {
		var info table.EventHandle
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		// info.AuthorCategory = 0
		result, err := table.InsertEventHandle(dbc, info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
	})

	router.POST("/eventHandle/user/add", func(c *gin.Context) {
		var info table.EventHandle
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		info.AuthorCategory = 0
		result, err := table.InsertEventHandle(dbc, info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
	})

	// 法官受理投诉 通知居民和PM
	router.POST("/eventHandle/court/accept", func(c *gin.Context) {
		var info table.EventHandle
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		info.AuthorCategory = 5 //法官
		result, err := table.InsertEventHandle(dbc, info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err = table.UpdateEventCourtAccepted(dbc, info.Index, 1)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})

		err = NoticePM(dbc, info.Index)
		if err != nil {
			glog.Error(err)
		}
		PushNotice2WX(info, dbc)
	})

	router.POST("/eventHandle/court/ask", func(c *gin.Context) {
		var info table.EventHandle
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		info.AuthorCategory = 5 //法官
		result, err := table.InsertEventHandle(dbc, info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
	})
}

//PushNotice2Court 发送微信通知给法官
func PushNotice2Court(courts []table.CourtWX, xqName string, eventIndex string) {
	for _, court := range courts {
		pjson := `{
			"touser": "` + court.OpenID + `",
			"template_id": "TdVxvtwH1i24ArEUcx1FGmWNFI_11WFZvDGfBJ9cjBw",
			"url": "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxa768bfacbb694944&redirect_uri=https%3A%2F%2Fwww.maszfglzx.com%2Fcomplaint-progress-court.html&response_type=code&scope=snsapi_base&state=` + eventIndex + `#wechat_redirect",
			"data": {
				"first": {
					"value": "法官您好，房管中心向您推送了` + xqName + `小区的投诉事件，信息如下："
				},
				"keyword1": {
					"value": "` + eventIndex + `"
				},
				"keyword2": {
					"value": "推送给法官"
				},
				"keyword3": {
					"value": "房管中心"
				},
				"remark": {
					"value": "请您及时进行处理"
				}
			}
		}`
		access_token := GetAccessToken()
		wxURL := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token
		PostJson(wxURL, pjson)
	}
}

//PushNotice2PM 通知物业公司人员
func PushNotice2PM(pmUsers []table.PMUser, xqName string, eventIndex string, pmName string) {
	for _, pmUser := range pmUsers {
		pjson := `{
			"touser": "` + pmUser.OpenID + `",
			"template_id": "TdVxvtwH1i24ArEUcx1FGmWNFI_11WFZvDGfBJ9cjBw",
			"url": "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxa768bfacbb694944&redirect_uri=https%3A%2F%2Fwww.maszfglzx.com%2Fcomplaint-progress-pm.html&response_type=code&scope=snsapi_base&state=` + eventIndex + `#wechat_redirect",
			"data": {
				"first": {
					"value": "您好，贵公司所服务的` + xqName + `有群众投诉，信息如下："
				},
				"keyword1": {
					"value": "` + eventIndex + `"
				},
				"keyword2": {
					"value": "已分派"
				},
				"keyword3": {
					"value": "` + pmName + `"
				},
				"remark": {
					"value": "请贵公司及时进行处理"
				}
			}
		}`
		access_token := GetAccessToken()
		wxURL := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token
		PostJson(wxURL, pjson)
	}
}

func GOVTalkPM(pmUsers []table.PMUser, xqName string, eventIndex string,
	pmName string, talkContent string) {
	for _, pmUser := range pmUsers {
		pjson := `{
			"touser": "` + pmUser.OpenID + `",
			"template_id": "TdVxvtwH1i24ArEUcx1FGmWNFI_11WFZvDGfBJ9cjBw",
			"url": "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxa768bfacbb694944&redirect_uri=https%3A%2F%2Fwww.maszfglzx.com%2Fcomplaint-progress-pm.html&response_type=code&scope=snsapi_base&state=` + eventIndex + `#wechat_redirect",
			"data": {
				"first": {
					"value": "您好，贵公司所服务的` + xqName + `有群众投诉，房管中心正在约谈，信息如下"
				},
				"keyword1": {
					"value": "` + eventIndex + `"
				},
				"keyword2": {
					"value": "约谈"
				},
				"keyword3": {
					"value": "` + pmName + `"
				},
				"remark": {
					"value": "约谈内容：` + talkContent + `"
				}
			}
		}`
		access_token := GetAccessToken()
		wxURL := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token
		PostJson(wxURL, pjson)
	}
}
