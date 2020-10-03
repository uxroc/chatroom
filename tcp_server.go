package main

import (
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	conn net.Conn
	name string
	writer *CommandWriter
}

type TCPServer struct {
	listener net.Listener
	clients []*client
	mutex *sync.Mutex
}

func NewServer() *TCPServer {
	return &TCPServer{
		mutex: &sync.Mutex{},
	}
}

func (s *TCPServer) Listen(addr string) (err error) {
	var listener net.Listener
	if listener, err = net.Listen("tcp", addr); err != nil {
		return
	}
	log.Printf("now listening on %v", addr)
	s.listener = listener
	return
}

func (s *TCPServer) Close() {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *TCPServer) Start() {
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Fatal(err)
		} else {
			client := s.accept(conn)
			go s.serve(client)
		}
	}
}

func (s *TCPServer) accept(conn net.Conn) *client {
	log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.clients) + 1)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	client := &client{
		conn: conn,
		name: conn.RemoteAddr().String(),
		writer: NewCommandWriter(conn),
	}

	s.clients = append(s.clients, client)

	return client
}

func (s *TCPServer) remove(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	found := false
	for i, c := range s.clients {
		if c == client {
			s.clients[i] = s.clients[len(s.clients) - 1]
			found = true
			break
		}
	}

	if found {
		s.clients = s.clients[:len(s.clients) - 1]
	}

	log.Printf("Closing connection from %v, total clients: %v", client.conn.RemoteAddr().String(), len(s.clients))
	if err := client.conn.Close(); err != nil {
		log.Fatalf("Error closing client: %v", err.Error())
	}
}

func (s *TCPServer) serve(client *client) {
	cmdReader := NewCommandReader(client.conn)
	defer s.remove(client)
	serving := true
	for serving {
		cmd, err := cmdReader.Read()
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Fatalf("Error reading command: %v", err.Error())
		}

		log.Printf("receving command: %v from client %v", cmd, client.conn.RemoteAddr())

		if cmd != nil {
			switch v := cmd.(type) {
			case *SendCommand:
				log.Printf("client %v sends command: %v", client.conn.RemoteAddr(), v)
				go s.Broadcast(&MessageCommand{
					Message: v.Message,
					Name: client.name,
				})
			case *NameCommand:
				log.Printf("client %v renamed to: %v", client.conn.RemoteAddr(), v.Name)
				client.name = v.Name
			case *ExitCommand:
				log.Printf("client %v exits", client.conn.RemoteAddr())
				serving = false
			default:
				log.Fatalf("Uknown command: %v", v)
			}
		}
	}
}

func (s *TCPServer) Broadcast(command Command) {
	log.Printf("Broading casting: %v", command)
	for _, client := range s.clients {
		client.writer.Write(command)
	}
}