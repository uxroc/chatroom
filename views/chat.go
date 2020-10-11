package views

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
	"time"
)

type post struct {
	username string
	msg string
}

type ChatView struct {
	tui.Box
	sendHandler func(string)
	historyView *tui.Box
}

func NewChatView() *ChatView {
	view := &ChatView{}

	view.historyView = tui.NewVBox()

	historyScroll := tui.NewScrollArea(view.historyView)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		view.sendHandler(e.Text())
		input.SetText("")
	})

	view.Append(chat)
	return view
}

func(chat *ChatView) OnSend(handler func(string)) {
	chat.sendHandler = handler
}

func(chat *ChatView) Welcome(username string) {
	chat.historyView.Append(tui.NewHBox(
		tui.NewLabel(time.Now().Format("15:04")),
		tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("Welcome %s!", username))),
		tui.NewSpacer(),
	))
}

func(chat *ChatView) AddMessage(username string, msg string) {
	chat.historyView.Append(tui.NewHBox(
		tui.NewLabel(time.Now().Format("15:04")),
		tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", username))),
		tui.NewLabel(msg),
		tui.NewSpacer(),
	))
}