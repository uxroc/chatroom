package main

import (
	"io"
	"log"
	"net"
)

type TCPChatClient struct {
	conn net.Conn
	cmdReader *CommandReader
	cmdWriter *CommandWriter
	name string
	incoming chan MessageCommand
	end chan bool
}

func (c *TCPChatClient) Close() {
	c.conn.Close()
	close(c.end)
}

func (c *TCPChatClient) Incoming() chan MessageCommand {
	return c.incoming
}

func NewClient() *TCPChatClient {
	return &TCPChatClient{
		incoming: make(chan MessageCommand),
		end: make(chan bool),
	}
}

func (c *TCPChatClient) Dial(addr string) (err error) {
	if c.conn, err = net.Dial("tcp", addr); err != nil {
		return
	}

	c.name = c.conn.LocalAddr().String()
	c.cmdReader = NewCommandReader(c.conn)
	c.cmdWriter = NewCommandWriter(c.conn)

	log.Printf("Dial succeed!")
	return
}

func (c *TCPChatClient) Send(cmd Command) error {
	//log.Printf("client command: %v", cmd)
	return c.cmdWriter.Write(cmd)
}

func (c *TCPChatClient) SendMessage(msg string) error {
	//log.Printf("client sends msg: %v", msg)
	return c.Send(&SendCommand{msg})
}

func (c *TCPChatClient) SetName(name string) error {
	//log.Printf("client sets name: %v", name)
	return c.Send(&NameCommand{name})
}

func (c *TCPChatClient) Start() {
	log.Println("client Starts...")
	for {
		select {
		case <-c.end:
			return
		default:
			cmd, err := c.cmdReader.Read()
			//log.Println(cmd)

			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Error reading command: %v", err.Error())
			}

			if cmd != nil {
				switch v := cmd.(type) {
				case *MessageCommand:
					c.incoming <- *v
				default:
					log.Fatalf("Unknown Command: %v", v)
				}
			}
		}
	}
}