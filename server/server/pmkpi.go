package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../db/table"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2"
)

func ExportPMKPIFile(pmkpis []table.PMKPI, fileName string, year int, quart int) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet(strconv.Itoa(year) + "年第" + strconv.Itoa(quart) +
		"季度物业考评")
	if err != nil {
		return err
	}
	row = sheet.AddRow()
	row.AddCell().Value = "年"
	row.AddCell().Value = "季度"
	row.AddCell().Value = "物业公司"
	row.AddCell().Value = "小区名"
	row.AddCell().Value = "黄色警告次数"
	row.AddCell().Value = "红色警告次数"
	row.AddCell().Value = "重大时间警告次数"
	row.AddCell().Value = "其他"
	row.AddCell().Value = "得分"
	for _, pmkpi := range pmkpis {
		row = sheet.AddRow()
		row.AddCell().Value = strconv.Itoa(pmkpi.Year)
		row.AddCell().Value = strconv.Itoa(pmkpi.Quarter)
		row.AddCell().Value = pmkpi.PMName
		row.AddCell().Value = pmkpi.XQName
		row.AddCell().Value = strconv.Itoa(pmkpi.YWNo)
		row.AddCell().Value = strconv.Itoa(pmkpi.RWNo)
		row.AddCell().Value = strconv.Itoa(pmkpi.IWNo)
		row.AddCell().Value = strconv.Itoa(pmkpi.Other)
		row.AddCell().Value = fmt.Sprintf("%.2f", pmkpi.Score)
	}
	err = file.Save(FileBasicPath + "/export/" + fileName)
	return err
}

func SetNameAndScore(pmkpi *table.PMKPI, pmName string, xqName string) {
	pmkpi.PMName = pmName
	pmkpi.XQName = xqName
	pmkpi.Score = 100.00 - float32(pmkpi.YWNo)*0.5 -
		float32(pmkpi.RWNo*1) - float32(pmkpi.IWNo*3) - float32(pmkpi.Other)
}

func startPMKPI(router *gin.RouterGroup, dbc *mgo.Database) {
	//按年季度导出所有小区的pmkpi
	router.POST("/pm/kpi/export", func(c *gin.Context) {
		var params QueryPMKPI
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		allXQs, err := table.FindAllXQs(dbc)
		var pmkpis []table.PMKPI
		for _, xq := range allXQs {
			err, pmkpi := table.FindPMKPI(dbc, xq.ID.Hex(), params.Year, params.Quarter)
			if err != nil {
				continue
			}
			pm := table.FindPMByKV(dbc, "xqid", xq.ID.Hex())
			xq := table.FindXQ(dbc, xq.ID.Hex())
			SetNameAndScore(&pmkpi, pm.Name, xq.Name)
			pmkpis = append(pmkpis, pmkpi)
		}
		fileName := "pmkpi_" + strconv.Itoa(params.Year) + strconv.Itoa(params.Quarter) + ".xlsx"
		err = ExportPMKPIFile(pmkpis, fileName, params.Year, params.Quarter)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"file": fileName}})
	})

	// 按小区ID、年、季度查询kpi
	router.POST("/pmkpi/query", func(c *gin.Context) {
		var params QueryPMKPI
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		err, pmkpi := table.FindPMKPI(dbc, params.XQID, params.Year, params.Quarter)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": pmkpi})
	})

	router.POST("/pmkpi", func(c *gin.Context) {
		var queryInfo QueryBasic
		err := c.BindJSON(&queryInfo)
		if err == nil {
			c.JSON(http.StatusOK, table.FindPMKPIsByName(dbc, queryInfo.Name,
				queryInfo.PageNo, queryInfo.PageSize))
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		}
	})

	router.POST("/pmkpi/kvs", func(c *gin.Context) {
		params := make(map[string]interface{})
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			log.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, table.FindPMKPIByKVs(dbc, params))
		}
	})

	router.POST("/pmkpi/kvs/page", func(c *gin.Context) {
		var params QueryKVs
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			c.JSON(http.StatusOK, table.FindPMKPIByKVsPage(dbc, params.KVs,
				params.PageNo, params.PageSize))
		}
	})

	router.POST("/pmkpi/update/other", func(c *gin.Context) {
		var pmkpi table.PMKPI
		err := c.BindJSON(&pmkpi)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
		} else {
			err = table.UpdatePMKPIOther(dbc, pmkpi)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"error": 0, "data": ""})
		}
	})
}
