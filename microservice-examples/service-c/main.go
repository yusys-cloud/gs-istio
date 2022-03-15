// Author: yangzq80@gmail.com
// Date: 2022-01-05
//
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func main() {

	router := gin.Default()

	msg := "service-c UP response ok "

	router.GET("/api/c", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			log.Infof("Header [%v]---[%v]", k, v)
		}

		c.JSON(http.StatusOK, msg)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":1003")
}
