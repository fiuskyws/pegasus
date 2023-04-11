package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewQueue(t *testing.T) {
	require := require.New(t)

	actual := New[int]()

	require.NotNil(actual)
	require.NotNil(actual._l)
	require.Nil(actual._l.Value)
	require.Nil(actual._l.Next)
}

func TestQueuePush(t *testing.T) {
	q := New[int]()

	t.Run("Check First Item", func(t *testing.T) {
		require := require.New(t)

		actual := q._l.Value

		require.Nil(actual)
	})

	t.Run("Push", func(t *testing.T) {
		require := require.New(t)

		expected := 10
		q.Push(expected)

		actual := *q._l.Value

		require.Equal(expected, actual)
	})
}

func TestQueueGet(t *testing.T) {
	q := New[int]()

	t.Run("Get Empty Queue", func(t *testing.T) {
		require := require.New(t)

		actual := q.Get()

		require.Nil(actual)
	})

	t.Run("Get w/ 1 value", func(t *testing.T) {
		require := require.New(t)

		expectedValue := 10
		q.Push(expectedValue)

		actualValue := q.Get()

		require.NotNil(actualValue)
		require.Equal(expectedValue, *actualValue)
	})
}

func TestQueuePop(t *testing.T) {
	q := New[int]()

	t.Run("Pop Empty Queue", func(t *testing.T) {
		require := require.New(t)

		actual := q.Pop()

		require.Nil(actual)
	})

	t.Run("Pop w/ 1 value", func(t *testing.T) {
		require := require.New(t)

		expectedValue := 10
		q.Push(expectedValue)

		actual := q.Pop()

		require.NotNil(actual)
	})
}
