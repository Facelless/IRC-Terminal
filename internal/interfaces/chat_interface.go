package interfaces

import (
	"remoteview/internal/domain"
)


type ChatService interface {
	Join(c *domain.Client)
	Quit(c *domain.Client)
	ChangeNick(c *domain.Client, newNick string)
	Broadcast(msg domain.Message)
	ListUsers() []string
}
