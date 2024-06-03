package main

import (
	"crypto/tls"
	"main/handler"
	// "crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"net"
	"time"
)

func main() {
	
	url := "newbinusmaya.binus.ac.id:443"


	conn, err := net.DialTimeout("tcp", url, 10*time.Second)
	handler.ErrorHand(err)

	defer conn.Close()

	tlsConn := tls.Client(conn, &tls.Config{
		InsecureSkipVerify: true, 
	})


	err = tlsConn.Handshake()
	handler.ErrorHand(err)


	state := tlsConn.ConnectionState()

	tlsVersion := getTLSVersionName(state.Version)


	cipherSuite := tls.CipherSuiteName(state.CipherSuite)

	if len(state.PeerCertificates) == 0 {
		fmt.Println("Sertifikat server yang ditemukan")
		return
	}
	cert := state.PeerCertificates[0]

	issuerOrg := getIssuerOrganization(cert.Issuer)

	fmt.Println("Versi TLS:", tlsVersion)
	fmt.Println("Cipher Suite:", cipherSuite)
	fmt.Println("Issuer Organization:", issuerOrg)
}


func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS13:
		return "TLS 1.3"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS10:
		return "TLS 1.0"
	default:
		return "TLS tidak diketahui"
	}
}


func getIssuerOrganization(issuer pkix.Name) string {
	if len(issuer.Organization) > 0 {
		return issuer.Organization[0]
	}
	return "Tidak ada informasi terkait"
}