package queue

type (
	// Queue is a linked-list data structure, that will be implemented
	// instead of a common `chan` to allow us to peek.
	// FIFO
	Queue[T any] struct {
		_l *list[T]
	}
)

// New returns a new instance of Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		_l: newList[T](),
	}
}

// Push - Add a item to the end of the queue.
func (q *Queue[T]) Push(v T) {
	q._l.Append(v)
}

// Get - Retrieves the Queue's first item, but
// doesn't delete it, if you call `Get()` again,
// it will return the same value as the first call.
func (q *Queue[T]) Get() *T {
	return q._l.Value
}

// Pop - Retrieves Queue's first item and
// removes it.
func (q *Queue[T]) Pop() *T {
	return q._l.Pop(0)
}
