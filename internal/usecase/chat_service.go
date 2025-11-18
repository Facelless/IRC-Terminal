package usecase

import (
	"remoteview/internal/domain"
	"remoteview/internal/interfaces"
	"sync"
)

type ChatService struct {
	mu sync.Mutex
	clients map[string]*domain.Client
}

func NewChatService() interfaces.ChatService {
	return &ChatService{clients: make(map[string]*domain.Client)}
}

func (s *ChatService) Join(c *domain.Client) {
	s.mu.Lock()
	s.clients[c.Nick] = c
	s.mu.Unlock()
	s.Broadcast(domain.Message{From: "@", Text: c.Nick + " entrou"})	
}

func (s *ChatService) Quit(c *domain.Client) {
	s.mu.Lock()
	delete(s.clients, c.Nick)
	s.mu.Unlock()
	s.Broadcast(domain.Message{From: "@", Text: c.Nick + " saiu"})
}

func (s *ChatService) ChangeNick(c *domain.Client, newNick string) {
	s.mu.Lock()
	delete(s.clients, c.Nick)
	old := c.Nick
	c.Nick = newNick
	s.clients[newNick] = c
	s.mu.Unlock()
	s.Broadcast(domain.Message{From: "@", Text: old + "agora e " + newNick})
}


func (s *ChatService) Broadcast(msg domain.Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, cl := range s.clients {
		cl.Conn.Write([]byte(msg.From + ": " + msg.Text + "\n"))
	}
}

func (s *ChatService) ListUsers() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := []string{}
	for k := range s.clients {
		list = append(list, k)
	}
	return list
}
