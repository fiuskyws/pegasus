package topic

import "github.com/google/uuid"

// TODO:
//   - Implement Queue data structure in Topic
//   - Add `Peek() T`  and `PeekAt(i uint) (*T)` methods
type (
	// Topic is a "named queue" that services can subscribe to.
	Topic[T any] struct {
		// q is the Topic's internal queue.
		q    chan T
		Name string
	}
)

// NewTopic creates a new Topic, if given name is empty, it generates
// an UUID as a name.
func NewTopic[T any](name string) (*Topic[T], error) {
	if name == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		name = id.String()
	}

	return &Topic[T]{
		q:    make(chan T, 10),
		Name: name,
	}, nil
}

// GetReceiver will return a read-only Queue.
func (t *Topic[T]) GetReceiver() <-chan T {
	return t.q
}

// GetSender will return a write-only Queue.
func (t *Topic[T]) GetSender() chan<- T {
	return t.q
}

// Pop will retrieve and delete a message from the output channel.
func (t *Topic[T]) Pop() (T, error) {
	msg := <-t.q
	return msg, nil
}

// Send will put a message in Topic input channel.
func (t *Topic[T]) Send(m T) error {
	t.q <- m
	return nil
}
