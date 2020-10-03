package main

import (
	"io"
)

type CommandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{writer}
}

func (w *CommandWriter) writeMsg(msg string) error {
	_, err := w.writer.Write([]byte(msg))

	return err
}

func (w *CommandWriter) Write(command Command) error {
	//log.Printf("writes: %v", command.toStr())
	err := w.writeMsg(command.toStr())

	return err
}
