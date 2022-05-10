// Author: yangzq80@gmail.com
// Date: 3/29/22
//
package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTCPGet(t *testing.T) {
	tcpAddress := "127.0.0.1:2001"
	tcpCon, err := net.Dial("tcp", tcpAddress)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(tcpCon)
}

func TestTcpEchoServer(t *testing.T) {
	// set up test parameters
	prefix := "hello"
	request := "world"
	want := "{\"code\""

	// start the TCP Echo Server
	os.Args = []string{"main", "9000,9001", prefix}
	go main()

	// wait for the TCP Echo Server to start
	time.Sleep(2 * time.Second)

	for _, addr := range []string{":2001"} {
		// connect to the TCP Echo Server
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatalf("couldn't connect to the server: %v", err)
		}
		defer conn.Close()

		// test the TCP Echo Server output
		if _, err := conn.Write([]byte(request + "\n")); err != nil {
			t.Fatalf("couldn't send request: %v", err)
		} else {
			reader := bufio.NewReader(conn)
			if response, err := reader.ReadBytes(byte('\n')); err != nil {
				t.Fatalf("couldn't read server response: %v", err)
			} else if !strings.HasPrefix(string(response), want) {
				t.Errorf("output doesn't match, wanted: %s, got: %s", want, response)
			}
		}
	}
}
