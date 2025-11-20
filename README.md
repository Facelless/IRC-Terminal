![share](https://i.postimg.cc/VkK78kMj/image.png)

# TermiChat

A lightweight, terminal-based IRC-style chat application written in Go, following Clean Architecture principles. Supports multiple clients, nick management, user lists. Designed for simplicity, extensibility, and easy terminal use.

---

## Features

- Terminal-based IRC-style chat
- Multiple clients support
- Nickname management (`/nick NEWNAME`)
- List connected users (`/who`)
- Quit command (`/quit`)
- Built with **Clean Architecture** in Go
- Easy to extend (channels, private messages, TLS, etc.)

---

## Folder Structure

```
chat-irc/
├── cmd/
│ ├── server/ # Server entry point
│ └── client/ # Client entry point
├── internal/
│ ├── domain/ # Core entities (Client, Message)
│ ├── usecase/ # Business logic (ChatService)
│ ├── infra/ # TCP server/client infrastructure
│ └── interfaces/ # Interface definitions (ports)
└── go.mod
  ```


---

## Installation

Make sure you have **Go 1.18+** installed.

Clone the repository:

```
git clone https://github.com/Facelless/IRC-Terminal.git
cd IRC-Terminal
```

## Running the Server
```
go run cmd/server/main.go
```

## Running the Client
```
go run cmd/client/main.go
```
