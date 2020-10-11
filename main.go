package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	args := os.Args[1:]
	switch args[0] {
	case "server":
		StartNewServer(args[1])
	case "client":
		StartNewClientUI(args[1])
	case "raw-client":
		StartNewClient(args[1])
	}
}
