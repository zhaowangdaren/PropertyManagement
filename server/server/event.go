package server

import (
	"fmt"
	"net/http"
	"strconv"

	"../db/table"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2"
)

func ExportEventOverviewFile(eventOverviews []table.EventOverview, streetName string, fileName string) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	file = xlsx.NewFile()
	fmt.Println("evenOverviews length", len(eventOverviews))
	sheet, err = file.AddSheet(strconv.Itoa(eventOverviews[0].Year) + "年")
	if err != nil {
		return err
	}
	row = sheet.AddRow()
	row.AddCell().Value = "年"
	row.AddCell().Value = "季度"
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
		row.AddCell().Value = strconv.Itoa(eventOverview.Quarter)
		row.AddCell().Value = strconv.Itoa(eventOverview.Month)
		row.AddCell().Value = eventOverview.StreetName
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

//QueryAllStreetEventOverviewByMonth
func QueryAllStreetEventOverviewByMonth(db *mgo.Database, year int, month int,
	userType int, allStreets []table.Street) ([]table.EventOverview, error) {
	// allStreets, err := table.FindAllStreets(db)
	var eventOverviews []table.EventOverview
	var err error
	for _, street := range allStreets {
		eventOverview, err := table.QueryEventOverview(db, street.ID.Hex(),
			year, month, userType)
		if err != nil {
			glog.Error(err.Error())
			continue
		}
		SetEventOverviewStreetName(&eventOverview, street.Name, month)
		eventOverviews = append(eventOverviews, eventOverview)
	}
	return eventOverviews, err
}

func SetEventOverviewStreetName(eventOverview *table.EventOverview, streetName string, month int) {
	eventOverview.StreetName = streetName
}

func ExportEventOverviewByStreet(dbc *mgo.Database, params QueryEventOverview) []table.EventOverview {
	var eventOverviews []table.EventOverview
	var err error
	var allStreets []table.Street
	if params.StreetID != "" {
		street, err := table.FindStreetByID(dbc, params.StreetID)
		if err != nil {
			return eventOverviews
		}
		allStreets = append(allStreets, street)
	} else {
		allStreets, err = table.FindAllStreets(dbc)
		if err != nil {
			return eventOverviews
		}
	}
	if params.Month > 0 { //查询某一月份所有街道
		eventOverviews, err = QueryAllStreetEventOverviewByMonth(dbc, params.Year,
			params.Month, params.UserType, allStreets)
	} else if params.Quarter > 0 && params.Quarter <= 4 {
		eventOverviews = ExportEventOverviewByQuarter(dbc, params, params.Quarter, allStreets)
	} else if params.Quarter == 0 {
		for quarter := 1; quarter <= 4; quarter++ {
			tmp := ExportEventOverviewByQuarter(dbc, params, quarter, allStreets)
			eventOverviews = append(eventOverviews, tmp...)
		}
	}
	if err != nil {
		glog.Error(err.Error())
	}
	return eventOverviews
}
func ExportEventOverviewByQuarter(dbc *mgo.Database, params QueryEventOverview, quarter int, streets []table.Street) []table.EventOverview {
	var eventOverviews []table.EventOverview
	eo1, err := QueryAllStreetEventOverviewByMonth(dbc, params.Year,
		quarter*3-2, params.UserType, streets)
	if err == nil {
		eventOverviews = append(eventOverviews, eo1...)
	}
	eo2, err := QueryAllStreetEventOverviewByMonth(dbc, params.Year,
		quarter*3-1, params.UserType, streets)
	if err == nil {
		eventOverviews = append(eventOverviews, eo2...)
	}
	eo3, err := QueryAllStreetEventOverviewByMonth(dbc, params.Year,
		quarter*3, params.UserType, streets)
	if err == nil {
		eventOverviews = append(eventOverviews, eo3...)
	}
	return eventOverviews
}
func getPageData(data []table.EventOverview, pageNo int, pageSize int) []table.EventOverview {
	sum := len(data)

	if (pageNo+1)*pageSize > sum {
		return data[(pageNo * pageSize):sum]
	}
	return data[(pageNo * pageSize):((pageNo + 1) * pageSize)]
}

func startEvent(router *gin.RouterGroup, dbc *mgo.Database) {
	//针对某一季度所有街道的数据
	router.POST("/event/overview/export/street", func(c *gin.Context) {
		var params QueryEventOverview
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		eventOverviews := ExportEventOverviewByStreet(dbc, params)
		fileName := "events_" + strconv.Itoa(params.Year) +
			strconv.Itoa(params.Quarter) + ".xlsx"
		err = ExportEventOverviewFile(eventOverviews, params.StreetName, fileName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"file": fileName}})
	})

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
		// result, err := table.QueryEventOverview(dbc, params.StreetID, params.Year,
		// 	params.Month, params.UserType)
		eventOverviews := ExportEventOverviewByStreet(dbc, params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		var result []table.EventOverview
		if params.PageSize == 0 {
			result = eventOverviews
		} else {
			result = getPageData(eventOverviews, params.PageNo, params.PageSize)
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"list": result, "sum": len(eventOverviews)}})
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
