// Author: yangzq80@gmail.com
// Date: 2022-01-05
//
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	SVC_C string = "http://service-c:2003"
)

func main() {

	router := gin.Default()

	msg := "service-b-v2 UP call service-c "

	router.GET("/api/b", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			log.Infof("Header [%v]---[%v]", k, v)
		}
		code, msg := callC(msg)
		c.JSON(code, msg)
	})

	// 访问超时
	router.GET("/api/timeout", func(c *gin.Context) {
		sec, _ := strconv.Atoi(c.DefaultQuery("second", "0"))
		startTime := time.Now()

		time.Sleep(time.Second * time.Duration(sec))

		log.Printf("Timeout cost:%v", time.Since(startTime).Seconds())

		c.JSON(http.StatusOK, fmt.Sprintf("cost %v", time.Since(startTime).Seconds()))
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":2002")
}

func callC(msg string) (int, string) {
	resp, err := http.Get(SVC_C + "/api/c")
	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, msg + "ERROR:" + err.Error()
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(msg)
		return http.StatusOK, msg + string(body)
	}
}
