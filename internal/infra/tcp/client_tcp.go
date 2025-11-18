package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient(addr string) {
	conn, _ := net.Dial("tcp", addr)
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Nick: ")
	nick, _ := r.ReadString('\n')
	conn.Write([]byte(nick))

	go func() {
		s := bufio.NewReader(conn)
		for {
			msg, err := s.ReadString('\n')
			if err != nil {
				os.Exit(0)
			}
			fmt.Print(msg)
		}
	}()

	for {
		text, _ := r.ReadString('\n')
		conn.Write([]byte(text))
		if text == "/quit\n" {
			return
		}
	}
}
