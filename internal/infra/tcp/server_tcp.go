package tcp

import (
	"bufio"
	"net"
	"remoteview/internal/domain"
	"remoteview/internal/interfaces"
	"strings"
)

type TCPServer struct {
	service interfaces.ChatService
}

func NewTCPServer(s interfaces.ChatService) *TCPServer {
	return &TCPServer{s}
}

func (t *TCPServer) Start() {
	ln, _ := net.Listen("tcp", "0.0.0.0:7777")
	for {
		conn, _ := ln.Accept()
		go t.handle(conn)
	}
}

func (t *TCPServer) handle(conn net.Conn) {
	r := bufio.NewReader(conn)
	nick, _ := r.ReadString('\n')
	c := &domain.Client{Conn: conn, Nick: strings.TrimSpace(nick)}
	t.service.Join(c)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			t.service.Quit(c)
			return
		}
		msg := strings.TrimSpace(line)

		if strings.HasPrefix(msg, "/nick ") {
			n := strings.TrimPrefix(msg, "/nick ")
			t.service.ChangeNick(c, n)
		} else if msg == "/who" {
			users := t.service.ListUsers()
			conn.Write([]byte(strings.Join(users, ", ") + "\n")) 
		} else if msg == "/quit" {
			t.service.Quit(c)
			return
		} else {
			t.service.Broadcast(domain.Message{From: c.Nick, Text: msg})
		}
	}
}
