package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
)

// Client holds info about connection
type Client struct {
	Id     int
	Conn   net.Conn
	Server *server
	// incoming chan string // Channel for incoming data from client
}

// TCP server
type server struct {
	// clients                  []*Client
	address string        // Address to open connection: localhost:9999
	replay  string        // Address to open connection: localhost:9999
	joins   chan net.Conn // Channel for new connections

	// dd       net.Conn
	incoming chan string // Channel for incoming data from client

	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message string)
}

// Read client data from channel
func (c *Client) listen() {
	reader := bufio.NewReader(c.Conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			c.Conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		c.Server.onNewMessage(c, message)
		go func() {
			c.Server.incoming <- message
		}()
	}
}

func (c *Client) Send(message string) error {
	_, err := c.Conn.Write([]byte(message + "\n"))
	return err
}

// Get Conn
func (c *Client) GetConn() net.Conn {
	return c.Conn
}

// Called right after server starts listening new client
func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *server) OnNewMessage(callback func(c *Client, message string)) {
	s.onNewMessage = callback
}

// Creates new Client instance and starts listening
func (s *server) newClient(Conn net.Conn) {
	client := &Client{
		Conn:   Conn,
		Server: s,
	}
	go client.listen()
	s.onNewClientCallback(client)
}

// Listens new connections channel and creating new client
func (s *server) listenChannels() {
	for {
		select {
		case Conn := <-s.joins:
			s.newClient(Conn)
		}
	}
}

// Creates new tcp server instance
func New(address string, replay string) *server {
	log.Println("Creating TCP :", address)
	server := &server{
		address:  address,
		replay:   replay,
		joins:    make(chan net.Conn),
		incoming: make(chan string),
	}
	// if replay != "" {
	// 	var err error
	// 	server.dd, err = net.Dial("tcp", replay)
	// 	if err == nil {
	// 		defer server.dd.Close()
	// 	} else {
	// 		fmt.Printf("DD_FAIL_IGNORE：%s\r\n", err.Error())
	// 	}
	// }
	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message string) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})
	go server.onSendDD()
	return server
}
func (s *server) sendDb(message string) (err error) {
	conn, err := net.Dial("tcp", s.replay)
	if err != nil {
		return errors.Wrap(err, "DD_LINK_FAILED")
	}
	defer conn.Close()
	_, err = conn.Write([]byte(message))
	if err != nil {
		return errors.Wrap(err, "DD_FAILED")
	}
	return
}
func (s *server) onSendDD() (err error) {
	if s.replay == "" {
		return
	}
	for {
		select {
		case msg := <-s.incoming:
			err = s.sendDb(msg)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
