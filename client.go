package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"main/handler"
	"net/http"
)

const (
  url = "https://localhost:8443"
)

func main() {
  cert, err := ioutil.ReadFile("ca.crt")


  handler.ErrorHand(err)


  caCertPool := x509.NewCertPool()
  caCertPool.AppendCertsFromPEM(cert)

  tlsConfig := &tls.Config{
    RootCAs: caCertPool,
  }

  tr := &http.Transport{
    TLSClientConfig: tlsConfig,
  }

  client := &http.Client{Transport: tr}
  resp, err := client.Get(url)

  handler.ErrorHand(err)

  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  
  handler.ErrorHand(err)

  log.Printf("Response body: %s\n", body)
}