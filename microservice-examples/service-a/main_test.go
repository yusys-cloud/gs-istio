// Author: yangzq80@gmail.com
// Date: 2022-01-05
//
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestSVCBTimeout(t *testing.T) {
	log.Printf("Call uri[%v] StatusCode[%v]",1,"qwqw")
	resp, err := http.Get("http://localhost:1001/api/timeout?second=5")
	if err != nil {
		log.Println(err.Error())
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}

}
