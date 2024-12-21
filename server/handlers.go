package server

import (
	"encoding/json"
	"fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const (
	// message types
	echo = "echo"
)

func (s *Server) EchoHandler(msg maelstrom.Message) error {
	body := make(map[string]any)
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return fmt.Errorf("umarshal message body: %w", err)
	}

	body["type"] = "echo_ok"

	return s.node.Reply(msg, body)
}
