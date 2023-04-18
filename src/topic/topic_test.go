package topic

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	uuidRegex = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
)

func TestTopic(t *testing.T) {
	t.Run("NewTopic", func(t *testing.T) {
		t.Run("With name", func(t *testing.T) {
			require := require.New(t)
			name := "topic"

			actual, err := NewTopic[string](name)

			require.Nil(err)
			require.NotNil(actual)
			require.Equal(name, actual.Name)
		})
		t.Run("Empty name", func(t *testing.T) {
			require := require.New(t)

			actual, err := NewTopic[string]("")

			require.Nil(err)
			require.NotNil(actual)
			require.Regexp(uuidRegex, actual.Name)
		})
	})

	t.Run("Send", func(t *testing.T) {
		require := require.New(t)
		topic, err := NewTopic[string]("")

		require.Nil(err)

		actual := topic.Send("foo")

		require.Nil(actual)
	})

	t.Run("Pop", func(t *testing.T) {
		require := require.New(t)
		topic, err := NewTopic[string]("")
		msg := "foo"

		require.Nil(err)

		{
			err := topic.Send(msg)
			require.Nil(err)
		}

		{
			actual, err := topic.Pop()
			require.Nil(err)
			require.Equal(msg, actual)
		}
	})
}
