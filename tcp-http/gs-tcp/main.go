// Author: yangzq80@gmail.com
// Date: 3/29/22
//
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	address := ":2001"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Printf("listen tcp on:%v\n", address)

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClientRequest(con)
	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)
			log.Printf("TCP.Accept() Reader from delim[\\n] ---> %s", clientRequest)
			if clientRequest == ":QUIT" {
				log.Println("client requested server to close the connection so closing")
				return
			}
		case io.EOF:
			log.Println("client closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		if _, err = con.Write(createResponse()); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}
	}
}

var t = 0

func createResponse() []byte {
	t++
	log.Printf("tcp response times:%v", t)
	return []byte(
		fmt.Sprintf(`{"code":%d,"data":{"id":"%d","created_at":"%s","des":"From tcp-server Write"}}`+"\n",
			http.StatusOK,
			t,
			time.Now().UTC().Format(time.RFC3339),
		),
	)
}
