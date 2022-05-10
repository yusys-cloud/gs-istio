// Author: yangzq80@gmail.com
// Date: 3/18/22
//
package server

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
)

const (
	checkHeader    = "x-ext-authz"
	allowedValue   = "allow"
	resultHeader   = "x-ext-authz-check-result"
	receivedHeader = "x-ext-authz-check-received"
	overrideHeader = "x-ext-authz-additional-header-override"
	resultAllowed  = "allowed"
	resultDenied   = "denied"
)

var (
	serviceAccount = flag.String("allow_service_account", "a", "allowed service account, matched against the service account in the source principal from the client certificate")
	httpPort       = flag.String("http", "8000", "HTTP server port")
	denyBody       = fmt.Sprintf("denied by ext_authz for not found header `%s: %s` in the request", checkHeader, allowedValue)
)

// ExtAuthzServer implements the ext_authz v2/v3 gRPC and HTTP check request API.
type ExtAuthzServer struct {
	httpServer *http.Server
	// For test only
	httpPort chan int
	grpcPort chan int
}

// ServeHTTP implements the HTTP check request.
func (s *ExtAuthzServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("[HTTP] read body failed: %v", err)
	}
	l := fmt.Sprintf("%s %s%s, headers: %v, body: [%s]\n", request.Method, request.Host, request.URL, request.Header, body)
	if allowedValue == request.Header.Get(checkHeader) {
		log.Printf("[HTTP][allowed]: %s", l)
		response.Header().Set(resultHeader, resultAllowed)
		response.Header().Set(overrideHeader, request.Header.Get(overrideHeader))
		response.Header().Set(receivedHeader, l)
		response.WriteHeader(http.StatusOK)
	} else {
		log.Printf("[HTTP][denied]: %s", l)
		response.Header().Set(resultHeader, resultDenied)
		response.Header().Set(overrideHeader, request.Header.Get(overrideHeader))
		response.Header().Set(receivedHeader, l)
		response.WriteHeader(http.StatusForbidden)
		response.Write([]byte(denyBody))
	}
}

func (s *ExtAuthzServer) startHTTP(address string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		log.Printf("Stopped HTTP server")
	}()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to create HTTP server: %v", err)
	}
	// Store the port for test only.
	s.httpPort <- listener.Addr().(*net.TCPAddr).Port
	s.httpServer = &http.Server{Handler: s}

	log.Printf("Starting HTTP server at %s", listener.Addr())
	if err := s.httpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func (s *ExtAuthzServer) Run(httpAddr string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go s.startHTTP(httpAddr, &wg)
	wg.Wait()
}

func (s *ExtAuthzServer) Stop() {
	log.Printf("GRPC server stopped")
	log.Printf("HTTP server stopped: %v", s.httpServer.Close())
}

func NewExtAuthzServer() *ExtAuthzServer {
	return &ExtAuthzServer{
		httpPort: make(chan int, 1),
	}
}
