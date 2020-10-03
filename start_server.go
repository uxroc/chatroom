package main

import "log"

func StartNewServer(addr string) {
	var server ChatServer
	server = NewServer()
	err := server.Listen(addr)
	if err != nil {
		log.Fatalf("Error listening to %v: %v", addr, err.Error())
		return
	}
	server.Start()
}
