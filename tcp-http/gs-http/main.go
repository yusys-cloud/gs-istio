// Author: yangzq80@gmail.com
// Date: 3/29/22
//
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func handleHttpIndex(w http.ResponseWriter, r *http.Request) {
	var err error
	tcpCon, err = net.Dial("tcp", tcpAddress)
	log.Printf("Dial to TCP [%v] success...", tcpAddress)
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpCon.Close()

	tcpToHttp(w, r, tcpCon)
	//for name, headers := range r.Header {
	//	for _, h := range headers {
	//		fmt.Fprintf(w, "hello: %v: %v\n", name, h)
	//	}
	//}
}

var tcpCon net.Conn

//var tcpAddress = "127.0.0.1:2001"
var tcpAddress = "gs-tcp:2001"

func main() {
	log.Println("Starting http [2002]")

	http.HandleFunc("/", handleHttpIndex)
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})
	err := http.ListenAndServe(":2002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func tcpToHttp(w http.ResponseWriter, r *http.Request, connection net.Conn) error {
	b, _ := ioutil.ReadAll(r.Body)
	request := bytes.NewBuffer(b)

	clientReader := bufio.NewReader(request)
	serverReader := bufio.NewReader(connection)

	for {
		req, err := clientReader.ReadString('\n')
		log.Println(req, err)
		connection.Write([]byte(strings.TrimSpace(req) + "\n"))

		res, err := serverReader.ReadString('\n')
		switch err {
		case nil:
			fmt.Fprintf(w, res)
			return nil
		case io.EOF:
			return errors.Wrap(err, "server closed the connection")
		default:
			return errors.Wrap(err, "server")
		}
	}
}
