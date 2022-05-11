// Author: yangzq80@gmail.com
// Date: 2022-01-05
//
package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	SVC_B string = "http://service-b:2002"
	//SVC_B string = "http://localhost:2002"
)

type Message struct {
	Name    string
	Content string
}

func main() {

	router := gin.Default()

	msg := "service-a UP call service-b "

	router.GET("/api/a", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			log.Infof("Header [%v]---[%v]", k, v)
		}
		log.Infof("Starting info request /api/a")
		log.Debugf("Starting debug request /api/a")
		log.Error("Starting error request /api/a")
		code, msg := callB(msg, SVC_B+"/api/b")
		c.JSON(code, msg)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// header 测试
	router.GET("/api/header", func(c *gin.Context) {
		code, msg := headerReq(msg, c.DefaultQuery("label-user", "groupA"))
		c.JSON(code, msg)
	})
	// 访问超时
	router.GET("/api/timeout", func(c *gin.Context) {
		code, msg := callB(msg, SVC_B+"/api/timeout?second="+c.DefaultQuery("second", "0"))
		c.JSON(code, msg)
	})
	// 无端口
	router.GET("/api/test1", func(c *gin.Context) {
		code, msg := callB(msg, "http://service-b/api/b")
		c.JSON(code, msg)
	})
	// 测试 ServiceEntry 外部接入 仅DNS可行 https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration
	router.GET("/api/se", func(c *gin.Context) {
		code, msg := callB(msg, "http://external-sc-svc-c:80/sc-service-c/api/sc/c")
		c.JSON(code, msg)
	})

	router.GET("/api/message/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, Message{"Join", "hello"})
	})
	router.GET("/api/message/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, Message{"Join", "hello"})
	})

	router.Run(":2001")
}

func callB(msg string, uri string) (int, string) {
	resp, err := http.Get(uri)
	log.Printf("Call uri[%v] StatusCode[%v]", uri, resp.StatusCode)
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

// /api/header?label-user=groupA 设置svcA->svcB请求header值("label-user":"groupA")
func headerReq(msg string, userLable string) (int, string) {
	req, _ := http.NewRequest("GET", SVC_B+"/api/b", nil)
	req.Header.Set("label-user", userLable)
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)

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
