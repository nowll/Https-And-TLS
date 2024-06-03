package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"main/handler"
)

const (
	port         = ":8443"
	responseBody = "Hello, TLS!"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")

	handler.ErrorHand(err)

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:      port,
		Handler:   router,
		TLSConfig: config,
	}

	log.Printf("Listening on %s...", port)

	err = server.ListenAndServeTLS("", "")

	handler.ErrorHand(err)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
}
