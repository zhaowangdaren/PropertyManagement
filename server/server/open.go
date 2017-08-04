//Package Server 公开API，例如供微信端用户投诉，查询街道
package server

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"../db"
	"../db/table"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

const FileBasicPath = "/Users/gtja/Documents/myDoc/golang/PropertyManagement/server/images/"

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

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		fileToken := fmt.Sprintf("%x", h.Sum(nil))
		err := c.SaveUploadedFile(file, FileBasicPath+fileToken)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": "上传失败"})
		} else {
			// md5Value, _ := Md5File(FileBasicPath + file.Filename)
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"md5": fileToken}})
		}
	})

	router.Static("/image", FileBasicPath)

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

	router.PUT("/event/update", func(c *gin.Context) {
		var info table.Event
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.UpdateEvent(dbc, info))
		}
	})
}
