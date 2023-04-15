package message

import (
	"fmt"
	"time"

	"github.com/fiuskyws/pegasus/src/proto"
)

type (
	// Message represents the data that's in the queue.
	Message struct {
		TopicName string `json:"topic_name"`
		// Timestamp registers the time the message was Published
		Timestamp *time.Time `json:"timestamp"`
		// Attr is a map of message Attributes.
		Attr map[string]any `json:"attr"`
		Body []byte         `json:"body"`
	}
)

const (
	errEmptyField = "field '%s' is empty"
)

func (m *Message) Validate() error {
	if m.TopicName == "" {
		return fmt.Errorf(errEmptyField, "topic_name")
	}
	return nil
}

func FromRequest(req *proto.SendRequest) (*Message, error) {
	msg := Message{
		Body:      []byte(req.Body),
		TopicName: req.TopicName,
	}

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	return &msg, nil
}
