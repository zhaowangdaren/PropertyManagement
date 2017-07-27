package server

import (
	"net/http"

	"../db/table"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func startEventHandle(router *gin.Engine, dbc *mgo.Database) {
	router.POST("/eventHandle/:index", func(c *gin.Context) {
		index := c.Param("index")
		c.JSON(http.StatusOK, table.FindEventHandle(dbc, index,
			1, 1))
	})
}
