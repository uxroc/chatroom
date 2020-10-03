package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func StartNewClient(addr string) {
	var client ChatClient
	client = NewClient()
	client.Dial(addr)

	go client.Start()

	//message receiver
	go func(incoming chan MessageCommand) {
		for {
			select {
			case v := <- incoming:
				fmt.Println(v.Name + ": " + v.Message)
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}(client.Incoming())

	//message sender
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	for scanner.Scan() && running {
		args := strings.Split(scanner.Text(), " ")
		switch strings.ToLower(args[0]) {
		case "msg", "m":
			if err := client.SendMessage(strings.Join(args[1:], " ")); err != nil {
				log.Fatalf("Error sending message: %v", err.Error())
			}
		case "name", "n":
			if err := client.SetName(strings.Join(args[1:], " ")); err != nil {
				log.Fatalf("Error naming user: %v", err.Error())
			}
		case "exit", "e":
			running = false
		default:
			log.Fatalf("Unknown command: %v", args[0])
		}
	}
	client.Close()
}
