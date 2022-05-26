package protocol

import "net"

type Client interface {
	Close()
}

type Protocol interface {
	NewClient(conn net.Conn) Client
	IOLoop(client Client) error
}
