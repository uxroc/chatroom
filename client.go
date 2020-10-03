package main

type ChatClient interface {
	Dial(addr string) error
	Send(cmd Command) error
	SendMessage(msg string) error
	SetName(name string) error
	Start()
	Close()
	Incoming() chan MessageCommand
}
