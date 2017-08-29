//Package Server 公开API，例如供微信端用户投诉，查询街道
package server

import (
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"../db/table"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
)

//FileBasicPath file basic path
const FileBasicPath = "/root/PropertyManagement/server/dist"
const WXConfPath = "/root/PropertyManagement/server/dist/wx.json"

//readFile readFile
func readFile(filename string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("readFile:", err.Error())
		return nil, err
	}
	var result = map[string]string{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println("Unmarshal:", err.Error())
	}
	return result, nil
}
func sendHttpRequest(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	defer response.Body.Close()
	var bodystr string
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr = string(body)
	} else {
		return ""
	}
	return bodystr
}

func getJson(url string) map[string]interface{} {
	bodystr := sendHttpRequest(url)
	var object interface{}
	err := json.Unmarshal([]byte(bodystr), &object)
	if err != nil {
		return nil
	} else {
		return object.(map[string]interface{})
	}
}

func Md5File(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func startOpen(router *gin.RouterGroup, dbc *mgo.Database) {
	flag.Parse()
	//查询街道信息，入参：name\pageNo\pageSize，当name为空时会查询所有的街道
	router.POST("/street", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindStreets(dbc, queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err)
		}
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

	//通过ID数组查找街道，返回的信息也是数组
	router.POST("/xq/ids", func(c *gin.Context) {
		var ids Values
		err := c.BindJSON(&ids)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindXQsByIDs(dbc, ids.Values))
		}
	})

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		fileToken := fmt.Sprintf("%x", h.Sum(nil))
		err := c.SaveUploadedFile(file, FileBasicPath+"/images/"+fileToken)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "上传失败"})
		} else {
			// md5Value, _ := Md5File(FileBasicPath + file.Filename)
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"md5": fileToken}})
		}
	})

	router.Static("/image", FileBasicPath+"/images/")

	router.POST("/event/add", func(c *gin.Context) {
		var info table.Event
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.InsertEvent(dbc, info))
		}
	})

	router.POST("/event/id/:id", func(c *gin.Context) {
		eventID := c.Param("id")
		c.JSON(http.StatusOK, table.FindEvents(dbc, eventID, 1, 1))
	})

	router.PUT("/event/addImgs", func(c *gin.Context) {
		var infos AddEventImg
		err := c.BindJSON(&infos)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			if len(infos.Imgs) > 9 {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": "图片数量超过9张啦>_<"})
			} else {
				c.JSON(http.StatusOK, table.AddEventImgs(dbc, infos.Index, infos.Imgs))
			}
		}
	})

	router.POST("/event/update", func(c *gin.Context) {
		var info table.Event
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.UpdateEvent(dbc, info))
		}
	})

	router.POST("/event/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindEventKVs(dbc, params))
		}
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

	router.POST("/pm/bind", func(c *gin.Context) {
		var info table.WXUser
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.InsertWXUser(dbc, info))
		}
	})

	wxInfo, _ := readFile("/root/PropertyManagement/server/dist/wx.json")
	router.POST("/wx/openid/:code", func(c *gin.Context) {
		code := c.Param("code")
		glog.Info("WeChat Get code:", code)
		glog.Info("WeChat appid appsecret", wxInfo)
		resp := getJson("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + wxInfo["appid"] + "&secret=" + wxInfo["appsecret"] + "&code=" + code + "&grant_type=authorization_code ")
		if resp == nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "获取用户信息失败"})
			return
		}
		glog.Info("WeChat Get openid", resp)
		glog.Flush()
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"openid": resp["openid"]}})
	})

	router.GET("/wx/event/types", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.FindEventTypes(dbc))
	})

	router.GET("/event/types/num", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.CountDiffTypeEvents(dbc))
	})

	glog.Flush()
}
