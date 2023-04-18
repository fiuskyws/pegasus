package message

import (
	"fmt"
	"time"
)

type (
	// Message represents the data that's in the queue.
	Message struct {
		TopicName string `json:"topic_name"`
		// Timestamp registers the time the message was Published
		Timestamp *time.Time `json:"timestamp"`
		// Attr is a map of message Attributes.
		Attr map[string]any `json:"attr"`
		Body string         `json:"body"`
	}

	requestAPI interface {
		GetBody() string
		GetTopicName() string
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

func FromRequest[T requestAPI](req T) (*Message, error) {
	msg := Message{
		Body:      req.GetBody(),
		TopicName: req.GetTopicName(),
	}

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	return &msg, nil
}
