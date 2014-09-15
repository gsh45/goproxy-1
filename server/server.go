package main

import (
	//"crypto/rand"
	"crypto/tls"
	//"crypto/x509"
	//"io/ioutil"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal("Error loading key pair: ", err)
	}
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, err := tls.Listen("tcp", ":8888", &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s\n", err)
		}
		defer conn.Close()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("server: accepted from %s\n", conn.RemoteAddr())
	buf := make([]byte, 512)
	//for {
		log.Printf("server: conn:waiting")
		n, err := conn.Read(buf)
        log.Printf("server: conn:read %d bytes\n", n)
		if err != nil {
			log.Printf("Server: conn read: %s\n", err)
			//break
		}
		//_, ok := conn.(*tls.Conn)
		//if ok {
			//state := tlsConn.ConnectionState()
			//sub := state.PeerCertificates[0].Subject
			//log.Println(tlsConn)
		//}

		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Printf("server: conn: write error\n", err)
			//break
		}
	//}
	log.Printf("server: conn closed\n")
}
