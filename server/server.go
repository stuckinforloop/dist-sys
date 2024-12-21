package server

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Server struct {
	node *maelstrom.Node
}

func New() *Server {
	n := maelstrom.NewNode()
	srv := &Server{
		n,
	}

	return srv
}

func (s *Server) RegisterHandlers() {
	s.node.Handle(echo, s.EchoHandler)
}

func (s *Server) Run() error {
	return s.node.Run()
}
