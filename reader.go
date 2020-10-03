package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	MESSAGE = "MESSAGE"
	SEND = "SEND"
	NAME = "NAME"
	EXIT = "EXIT"
)

type CommandReader struct {
	reader *bufio.Reader
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{bufio.NewReader(reader)}
}

func (r *CommandReader) Read() (c Command, err error) {
	var line string
	if line, err = r.reader.ReadString('\n'); err != nil {
		return
	}
	arr := strings.Split(line[:len(line) - 1], " ")
	switch arr[0] {
	case MESSAGE:
		return &MessageCommand{arr[1], strings.Join(arr[2:], " ")}, nil
	case SEND:
		return &SendCommand{strings.Join(arr[1:], " ")}, nil
	case NAME:
		return &NameCommand{strings.Join(arr[1:], " ")}, nil
	case EXIT:
		return &ExitCommand{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("UnknownCommand: %v", arr[0]))
	}
}
