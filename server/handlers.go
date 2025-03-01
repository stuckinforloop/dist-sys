package server

import (
	"encoding/json"
	"fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const (
	// message types
	echo     = "echo"
	generate = "generate"
)

func (s *Server) EchoHandler(msg maelstrom.Message) error {
	body, err := parseBody(msg.Body)
	if err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	body["type"] = returnType(echo)

	return s.node.Reply(msg, body)
}

func (s *Server) GenerateHandler(msg maelstrom.Message) error {
	body, err := parseBody(msg.Body)
	if err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	return s.node.Reply(msg, body)
}

func parseBody(msgBody json.RawMessage) (map[string]any, error) {
	body := make(map[string]any)
	if err := json.Unmarshal(msgBody, &body); err != nil {
		return nil, fmt.Errorf("umarshal message body: %w", err)
	}

	return body, nil
}

func returnType(s string) string {
	return s + "_" + "ok"
}
