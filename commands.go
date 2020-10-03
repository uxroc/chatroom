package main

import "fmt"

type Command interface {
	toStr() string
}

type SendCommand struct {
	Message string
}

type NameCommand struct {
	Name string
}

type MessageCommand struct {
	Name    string
	Message string
}

type ExitCommand struct {
	Name string

}

func (ec *ExitCommand) toStr() string {
	return fmt.Sprintf("EXIT %v\n", ec.Name)
}

func (sc *SendCommand) toStr() string {
	return fmt.Sprintf("SEND %v\n", sc.Message)
}

func (nc *NameCommand) toStr() string {
	return fmt.Sprintf("NAME %v\n", nc.Name)
}

func (mc *MessageCommand) toStr() string {
	return fmt.Sprintf("MESSAGE %v %v\n", mc.Name, mc.Message)
}
