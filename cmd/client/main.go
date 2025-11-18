package main

import "remoteview/internal/infra/tcp"

func main() {
	tcp.StartClient("127.0.0.1:7777")
}
