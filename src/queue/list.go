package queue

// TODO:
//   - Improve current implementation, to be memory/concurrency/parallel safe.

type (
	// list is a linked-list data structure implementation.
	list[T any] struct {
		Value *T
		Next  *list[T]
	}
)

// newList returns a new list with given type T.
func newList[T any]() *list[T] {
	return &list[T]{}
}

// Append will Append a Value.
func (l *list[T]) Append(v T) {
	if l.Value == nil {
		l.Value = &v
		return
	}

	l.Next = &list[T]{
		Value: &v,
	}
}

// Empty returns whether the list is empty or not.
func (l *list[T]) Empty() bool {
	return (l.Value == nil) && (l.Next == nil)
}

// Head returns list's head.
func (l *list[T]) Head() *T {
	return l.Value
}

func (l *list[T]) Last() *T {
	current := l
	for {
		if current.Next == nil {
			return current.Value
		}
		current = current.Next
	}
}

// Len return the amount of item in the list.
func (l *list[T]) Len() int {
	i := 0
	current := l
	for {
		if current.Next == nil {
			return i
		}
		i++
		current = current.Next
	}
}

// ListAt - returns the list instance at position `p`.
func (l *list[T]) ListAt(p int) *list[T] {
	current := *l
	if l.Next == nil && l.Value == nil {
		return nil
	}
	for i := 0; i <= p; i++ {
		if i == p {
			return current.Next
		}
		if current.Next == nil {
			break
		}
		current = *current.Next
	}
	return nil
}

// Pop - removes a value from the Linked list.
func (l *list[T]) Pop(p int) *T {
	current := *l
	currentPtr := l
	if l.Next == nil && l.Value == nil {
		return nil
	}
	for i := 0; i <= p; i++ {
		if i == p {
			val := *current.Value
			if currentPtr.Next != nil {
				currentPtr.Value = currentPtr.Next.Value
				currentPtr.Next = currentPtr.Next.Next
			} else {
				currentPtr.Value = nil
			}
			return &val
		}
		if current.Next == nil {
			break
		}
		current = *current.Next
		currentPtr = currentPtr.Next
	}
	return nil
}

// Prepend will add a Value to end of the list.
func (l *list[T]) Prepend(v T) {
	currentlist := *l
	newL := newList[T]()

	newL.Value = &v
	newL.Next = &currentlist

	*l = *newL
}

// Tail - returns the list without the first item.
func (l *list[T]) Tail() *list[T] {
	return l.Next
}

// ValueAt - Returns a Value at position p.
func (l *list[T]) ValueAt(p int) *T {
	current := *l
	if l.Value == nil {
		return nil
	}
	for i := 0; i <= p; i++ {
		if i == p {
			return current.Value
		}
		if current.Next == nil {
			break
		}
		current = *current.Next
	}
	return nil
}
