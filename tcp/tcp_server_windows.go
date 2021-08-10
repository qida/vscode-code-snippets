package tcp

import (
	"crypto/tls"
	"log"
	"net"
)

// // Start network server
// func (s *server) Listen() {
// 	go s.listenChannels()

// 	listener, err := net.Listen("tcp", s.address)
// 	if err != nil {
// 		log.Fatal("Error starting TCP server")
// 	}
// 	defer listener.Close()

// 	for {
// 		conn, _ := listener.Accept()
// 		s.joins <- conn
// 	}
// }

// Listen starts network server
func (s *server) Listen() {
	var listener net.Listener
	var err error
	if s.config == nil {
		listener, err = net.Listen("tcp", s.address)
	} else {
		listener, err = tls.Listen("tcp", s.address, s.config)
	}
	if err != nil {
		log.Fatal("Error starting TCP server.\r\n", err)
	}
	defer listener.Close()
	var relay net.Conn
	if s.relay != "" {
		relay, err = net.Dial("tcp", s.relay)
		if err != nil {
			log.Fatal("Error starting TCP relay.\r\n", err)
		}
		defer relay.Close()
	}
	for {
		conn, _ := listener.Accept()
		client := &Client{
			conn:   conn,
			relay:  relay,
			Server: s,
		}
		go client.listen()
	}
}
