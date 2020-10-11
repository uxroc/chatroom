package views

import (
	"github.com/marcusolsson/tui-go"
)

type LoginView struct {
	tui.Box
	loginHandler func(string)
}

func NewLoginView() *LoginView {
	login := &LoginView{}

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(tui.NewLabel("Please enter your name: "), input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		tui.NewPadder(10, 0, inputBox),
		tui.NewSpacer(),
	)

	input.OnSubmit(func(e *tui.Entry){
		login.loginHandler(e.Text())
	})

	login.Append(wrapper)

	return login
}

func (login *LoginView) OnLogin(handler func(string)) {
	login.loginHandler = handler
}