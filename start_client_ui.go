package main

import (
	"./views"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"log"
)

func initClient(addr string) (ChatClient, <-chan error) {
	var client ChatClient
	client = NewClient()
	if err := client.Dial(addr); err != nil {
		log.Fatalf("Error dialing %v: %v", addr, err.Error())
	}

	errChan := make(chan error)
	go func(errChan chan error) {
		errChan <- client.Start()
	}(errChan)

	return client, errChan
}

func StartNewClientUI(addr string) {
	client, errChan := initClient(addr)

	login := views.NewLoginView()
	chat := views.NewChatView()

	ui, err := tui.New(login)
	if err != nil {
		log.Fatal(err.Error())
	}

	chat.OnSend(func(msg string) {
		client.SendMessage(msg)
	})

	login.OnLogin(func(username string) {
		client.SetName(username)
		client.SendMessage(fmt.Sprintf("%s joined chat!", username))
		ui.SetWidget(chat)
		chat.Welcome(username)
	})

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Ctrl+c", func() { ui.Quit() })

	go func(incoming chan MessageCommand) {
		for {
			select {
			case v := <- incoming:
				ui.Update(func() {
					chat.AddMessage(v.Name, v.Message)
				})
			}
		}
	}(client.Incoming())

	go func(ui tui.UI, errChan <-chan error) {
		e := <-errChan
		ui.Quit()
		log.Fatalf("Error: %s", e.Error())
	}(ui, errChan)

	if err := ui.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
