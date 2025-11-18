package domain

import "net"

type Client struct {
	Conn net.Conn
	Nick string
}
