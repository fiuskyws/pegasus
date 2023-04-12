package message

import "time"

type (
	// Message represents the data that's in the queue.
	Message struct {
		// Timestamp registers the time the message was Published
		Timestamp *time.Time `json:"timestamp"`
		// Attr is a map of message Attributes.
		Attr map[string]any `json:"attr"`
		Body []byte         `json:"body"`
	}
)
