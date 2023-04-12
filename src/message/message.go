package message

import "time"

type (
	// Message represents the data that's in the queue.
	Message struct {
		// Timestamp registers the time the message was Published
		Timestamp *time.Time `json:"timestamp"`
		Body      []byte     `json:"body"`
	}
)

func (m Message) BodyBytes() []byte {
	return m.Body
}

func (m Message) BodyString() string {
	return string(m.Body)
}
