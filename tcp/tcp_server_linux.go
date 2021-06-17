package tcp

import (
	"log"
	"net"
	"time"

	"github.com/felixge/tcpkeepalive"
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
		kaConn, _ := tcpkeepalive.EnableKeepAlive(conn)
		kaConn.SetKeepAliveIdle(30 * time.Second)
		kaConn.SetKeepAliveCount(4)
		kaConn.SetKeepAliveInterval(5 * time.Second)
		s.joins <- kaConn
	}
}
