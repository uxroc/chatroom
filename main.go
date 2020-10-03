package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	args := os.Args[1:]
	if args[0] == "server" {
		StartNewServer(args[1])
	} else {
		StartNewClient(args[1])
	}
}
