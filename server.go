package main

import(
	"fmt"
	"net"
	"time"
	"log"
	"io"
	"crypto/tls"
	"crypto/x509"
	
)



//listen on port 9999
func main() {
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	service := "127.0.0.1:9999"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
//TLS handshake


//obtain requestFromClient


//find requestedMethod in GreenPagesDirectory
// lookups in id / requestedMethod Map
// provides host/port of provider from unique id

//


// Provider registers methods to GreenPagesDirectory
// server assigns unique id to host/port



//
