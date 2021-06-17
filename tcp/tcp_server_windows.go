package tcp

import (
	"log"
	"net"
)

// Start network server
func (s *server) Listen() {
	go s.listenChannels()

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		s.joins <- conn
	}
}
