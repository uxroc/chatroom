package main

import (
	"./views"
	"github.com/marcusolsson/tui-go"
	"log"
)

func initClient(addr string) ChatClient {
	var client ChatClient
	client = NewClient()
	if err := client.Dial(addr); err != nil {
		log.Fatalf("Error dialing %v: %v", addr, err.Error())
	}

	go client.Start()

	return client
}

func StartNewClientUI(addr string) {
	client := initClient(addr)
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

	if err := ui.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
