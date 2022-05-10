// Author: yangzq80@gmail.com
// Date: 3/18/22
//
package main

import (
	"flag"
	"github.com/yusys-cloud/istio-extauth/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()
	s := server.NewExtAuthzServer()
	go s.Run(":8000")
	defer s.Stop()

	// Wait for the process to be shutdown.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
