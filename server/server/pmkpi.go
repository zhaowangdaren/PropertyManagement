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
	row.AddCell().Value = "街道"
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
		row.AddCell().Value = pmkpi.StreetName
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

func SetNameAndScore(pmkpi *table.PMKPI, streetName string, pmName string, xqName string) {
	pmkpi.StreetName = streetName
	pmkpi.PMName = pmName
	pmkpi.XQName = xqName
	pmkpi.Score = 100.00 - float32(pmkpi.YWNo)*0.5 -
		float32(pmkpi.RWNo*1) - float32(pmkpi.IWNo*3) - float32(pmkpi.Other)
}

//ExportPMKPIBYStreet 按街道导出PMKPI
func ExportPMKPI(router *gin.RouterGroup, db *mgo.Database) {
	router.POST("/pm/kpi/export/street", func(c *gin.Context) {
		var params QueryPMKPI
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		var allStreets []table.Street
		if params.StreetID == "" {
			allStreets, err = table.FindAllStreets(db)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
				return
			}
		} else {
			street, err := table.FindStreetByID(db, params.StreetID)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
				return
			}
			allStreets = append(allStreets, street)
		}
		var pmkpis []table.PMKPI
		for _, street := range allStreets {
			tempPmkpi, terr := FindPMKPIByStreet(db, params.Year, params.Quarter, street)
			if terr != nil {
				continue
			}
			pmkpis = append(pmkpis, tempPmkpi...)
		}
		fileName := "pmkpi_" + params.StreetID + "_" + strconv.Itoa(params.Year) + "_" + strconv.Itoa(params.Quarter) + ".xlsx"
		err = ExportPMKPIFile(pmkpis, fileName, params.Year, params.Quarter)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"file": fileName}})
	})
}

func FindPMKPIByStreet(db *mgo.Database, year int, quarter int, street table.Street) ([]table.PMKPI, error) {
	var result []table.PMKPI
	xqs, err := table.FindXQsByStreetID(db, street.ID.Hex())
	if err != nil {
		return result, err
	}
	fmt.Println("xqs", xqs)
	for _, xq := range xqs {
		err, pmkpi := table.FindPMKPI(db, xq.ID.Hex(), year, quarter)
		if err != nil {
			continue
		}
		pm := table.FindPMByKV(db, "xqid", xq.ID.Hex())
		SetNameAndScore(&pmkpi, street.Name, pm.Name, xq.Name)
		result = append(result, pmkpi)
	}
	return result, err
}

func getPMKPIPageData(array []table.PMKPI, pageNo int, pageSize int) []table.PMKPI {
	sum := len(array)

	if (pageNo+1)*pageSize > sum {
		return array[(pageNo * pageSize):sum]
	}
	return array[(pageNo * pageSize):((pageNo + 1) * pageSize)]
}

func startPMKPI(router *gin.RouterGroup, dbc *mgo.Database) {
	//按年季度导出所有小区的pmkpi
	ExportPMKPI(router, dbc)
	// 按街道、年、季度查询kpi
	router.POST("/pmkpi/query", func(c *gin.Context) {
		var params QueryPMKPI
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		var allStreets []table.Street
		if params.StreetID == "" {
			allStreets, err = table.FindAllStreets(dbc)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
				return
			}
		} else {
			street, err := table.FindStreetByID(dbc, params.StreetID)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
				return
			}
			allStreets = append(allStreets, street)
		}
		var pmkpis []table.PMKPI
		for _, street := range allStreets {
			tempPmkpi, terr := FindPMKPIByStreet(dbc, params.Year, params.Quarter, street)
			if terr != nil {
				continue
			}
			pmkpis = append(pmkpis, tempPmkpi...)
		}

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": 1, "data": err.Error()})
			return
		}
		var result []table.PMKPI
		if params.PageSize == 0 {
			result = pmkpis
		} else {
			result = getPMKPIPageData(pmkpis, params.PageNo, params.PageSize)
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "data": gin.H{"list": result, "sum": len(pmkpis)}})
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
