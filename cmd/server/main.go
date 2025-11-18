package main

import (
	"remoteview/internal/infra/tcp"
	"remoteview/internal/usecase"
)

func main() {
	s := usecase.NewChatService()
	server := tcp.NewTCPServer(s)
	server.Start()
}
