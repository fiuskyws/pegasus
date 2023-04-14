package manager

import (
	"testing"

	"github.com/fiuskyws/pegasus/src/message"
	"github.com/stretchr/testify/require"
)

func TestNewManager(t *testing.T) {
	require := require.New(t)

	actual := NewManager()

	require.NotNil(actual)

}

func TestNewTopic(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()
		actualErr := m.NewTopic("")

		require.Nil(actualErr)
	})

	t.Run("Topic Exists", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		{
			actual := m.NewTopic("topic")
			require.Nil(actual)
		}
		{
			actual := m.NewTopic("topic")
			require.NotNil(actual)
		}
	})
}

func TestGetTopicNames(t *testing.T) {
	t.Run("Success - Empty", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		actual := m.GetTopicNames()

		require.Empty(actual)
	})

	t.Run("Success - 1 Topic", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		{
			actual := m.NewTopic("topic")
			require.Nil(actual)
		}

		{
			expected := []string{"topic"}
			actual := m.GetTopicNames()

			require.Equal(expected, actual)
		}
	})

	t.Run("Success - 10 Topics", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		amountOfTopics := 10

		{
			for i := 0; i < amountOfTopics; i++ {
				actual := m.NewTopic("")
				require.Nil(actual)
			}
		}

		{
			topics := m.GetTopicNames()

			actual := len(topics)
			require.Equal(amountOfTopics, actual)
		}
	})
}

func TestSend(t *testing.T) {
	t.Run("Invalid Message", func(t *testing.T) {
		t.Run("Empty TopicName", func(t *testing.T) {
			require := require.New(t)
			m := NewManager()

			actual := m.Send(&message.Message{
				Body: []byte("body"),
			})

			require.NotNil(actual)
		})

		t.Run("Empty Body", func(t *testing.T) {
			require := require.New(t)
			m := NewManager()

			actual := m.Send(&message.Message{
				TopicName: "topic",
			})

			require.NotNil(actual)
		})
	})

	t.Run("Topic Doesn't Exists", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		actual := m.Send(&message.Message{
			TopicName: "some topic",
			Body:      []byte("body"),
		})

		require.NotNil(actual)
	})

	t.Run("Topic Doesn't Exists", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		{
			err := m.NewTopic("topic")
			require.Nil(err)
		}

		{
			actual := m.Send(&message.Message{
				TopicName: "wrong_topic",
				Body:      []byte("body"),
			})

			require.NotNil(actual)
		}
	})

	t.Run("Valid", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		{
			err := m.NewTopic("topic")
			require.Nil(err)
		}

		{
			actual := m.Send(&message.Message{
				TopicName: "topic",
				Body:      []byte("body"),
			})

			require.Nil(actual)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Topic Doesn't Exists", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		actual, err := m.Pop("some_topic")

		require.NotNil(err)
		require.Nil(actual)
	})

	t.Run("Topic Doesn't Exists", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		{
			err := m.NewTopic("topic")
			require.Nil(err)
		}

		{
			actual, err := m.Pop("some_topic")

			require.NotNil(err)
			require.Nil(actual)
		}
	})

	t.Run("Valid", func(t *testing.T) {
		require := require.New(t)
		m := NewManager()

		{
			err := m.NewTopic("topic")
			require.Nil(err)
		}

		{
			err := m.Send(&message.Message{
				TopicName: "topic",
				Body:      []byte("body"),
			})

			require.Nil(err)
		}

		{
			actual, err := m.Pop("topic")

			require.Nil(err)
			require.NotNil(actual)
		}
	})
}
