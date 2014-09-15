package client

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io/ioutil"
)

func main() {
	publicKey, _ := ioutil.ReadFile("cert.pem")
	privateKey, _ := ioutil.ReadFile("key.pem")

	cert := tls.Certificate{Certificate: [][]byte{publicKey}, PrivateKey: privateKey}
}
