package server

import (
	"fmt"
	"net"
)

// Maybe delegate the connection off onto a structure that knows how
// to deal all things regarding the message

// There probably should be some parsing of the message and once the message is
// parsed, then we should serve back to client an appropriate action
func serveConnection(conn net.Conn) error {
	p := make([]byte, 4)
	_, err := conn.Read(p)
	if err != nil {
		return err
	}

	fmt.Printf("This is the message %s\n", string(p))
	return nil
}

func getListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func Run() error {
	l, err := getListener()
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		serveConnection(conn)
	}
}
