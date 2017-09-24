package server

import (
	"fmt"
	"net/http"
	"strconv"

	"../db/table"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2"
)

func ExportEventOverviewFile(eventOverviews []table.EventOverview, streetName string, fileName string) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet(strconv.Itoa(eventOverviews[0].Year) + "年" +
		eventOverviews[0].StreetName + "投诉数据")
	row = sheet.AddRow()
	row.AddCell().Value = "年"
	row.AddCell().Value = "月份"
	row.AddCell().Value = "街道名"
	row.AddCell().Value = "总投诉"
	row.AddCell().Value = "未处理"
	row.AddCell().Value = "未处理占比"
	row.AddCell().Value = "已处理"
	row.AddCell().Value = "已处理占比"
	row.AddCell().Value = "未解决"
	row.AddCell().Value = "未解决占比"
	row.AddCell().Value = "已解决"
	row.AddCell().Value = "已解决占比"
	for _, eventOverview := range eventOverviews {
		row = sheet.AddRow()
		row.AddCell().Value = strconv.Itoa(eventOverview.Year)
		row.AddCell().Value = strconv.Itoa(eventOverview.Month)
		row.AddCell().Value = streetName
		row.AddCell().Value = strconv.Itoa(eventOverview.Sum)
		row.AddCell().Value = strconv.Itoa(eventOverview.Unhandle)
		if eventOverview.Sum == 0 {
			row.AddCell().Value = "0%"
		} else {
			row.AddCell().Value = fmt.Sprintf("%.2f", (float32(eventOverview.Unhandle)/
				float32(eventOverview.Sum))*100) + "%"
		}
		row.AddCell().Value = strconv.Itoa(eventOverview.Handled)
		if eventOverview.Sum == 0 {
			row.AddCell().Value = "0%"
		} else {
			row.AddCell().Value = fmt.Sprintf("%.2f", (float32(eventOverview.Handled)/
				float32(eventOverview.Sum))*100) + "%"
		}
		row.AddCell().Value = strconv.Itoa(eventOverview.Unsolved)
		if eventOverview.Sum == 0 {
			row.AddCell().Value = "0%"
		} else {
			row.AddCell().Value = fmt.Sprintf("%.2f", (float32(eventOverview.Unsolved)/
				float32(eventOverview.Sum))*100) + "%"
		}
		row.AddCell().Value = strconv.Itoa(eventOverview.Solved)
		if eventOverview.Sum == 0 {
			row.AddCell().Value = "0%"
		} else {
			row.AddCell().Value = strconv.Itoa((eventOverview.Solved/
				eventOverview.Sum)*100) + "%"
		}
	}
	err = file.Save(FileBasicPath + "/export/" + fileName)
	return err
}
func startEvent(router *gin.RouterGroup, dbc *mgo.Database) {
	//TODO
	router.POST("/event/overview/export", func(c *gin.Context) {
		var params QueryEventOverview
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		var eventOverviews []table.EventOverview
		for month := 1; month <= 12; month++ {
			eventOverview, err := table.QueryEventOverview(dbc, params.StreetID,
				params.Year, month, params.UserType)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
				return
			}
			temp := eventOverview
			eventOverviews = append(eventOverviews, temp)
		}
		fileName := strconv.Itoa(params.Year) + params.StreetID + ".xlsx"

		err = ExportEventOverviewFile(eventOverviews, params.StreetName, fileName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"file": fileName}})
	})

	router.POST("/event/overview", func(c *gin.Context) {
		var params QueryEventOverview
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
		result, err := table.QueryEventOverview(dbc, params.StreetID, params.Year,
			params.Month, params.UserType)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": result})
	})

	router.POST("/event", func(c *gin.Context) {
		var queryInfo QueryEvent
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindEvents(dbc, queryInfo.Index,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.GET("/event/del/index/:index", func(c *gin.Context) {
		index := c.Param("index")
		err := table.DelEvent(dbc, index)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
		}
	})

	router.POST("/event/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventKVs(dbc, params))
		}
	})

	router.POST("/event/kvs/page", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventKVsPage(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})

	router.POST("/event/kvs/page/gov", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindEventKVsPageForGov(dbc, params.KVs, params.PageNo, params.PageSize))
		}
	})

	// 查询当个kv，并分页显示
	router.POST("/event/kv", func(c *gin.Context) {
		var params QueryKV
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK,
				table.FindEventKV(dbc,
					params.Key, params.Value, params.PageNo, params.PageSize))
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

	router.POST("/event/type/add", func(c *gin.Context) {
		var info table.EventType
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.InsertEventType(dbc, info.Type))
		}
	})
	//删除事件类型
	router.POST("/event/type/del", func(c *gin.Context) {
		var info table.EventType
		err := c.BindJSON(&info)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.DelEventType(dbc, info.Type))
		}
	})

	router.GET("/event/key/nums/:key", func(c *gin.Context) {
		key := c.Param("key")
		maxRecs, err := strconv.Atoi(c.Query("recs"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, table.CountDiffKeyEvents(dbc, key, maxRecs))
	})

	router.GET("/event/today/events", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.FindTodayEvents(dbc))
	})

	router.GET("/events/num", func(c *gin.Context) {
		c.JSON(http.StatusOK, table.CountEvents(dbc))
	})
}
