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
	// echo workload
	s.node.Handle(echo, s.EchoHandler)

	// unique-ids workload
	s.node.Handle(generate, s.GenerateHandler)
}

func (s *Server) Run() error {
	return s.node.Run()
}
