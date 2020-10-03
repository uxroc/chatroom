package main

type ChatServer interface {
	Listen(addr string) error
	Broadcast(command Command)
	Start()
	Close()
}
