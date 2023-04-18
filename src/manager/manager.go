package manager

import (
	"fmt"

	"github.com/fiuskyws/pegasus/src/message"
	"github.com/fiuskyws/pegasus/src/topic"
)

type (
	// Manager handles all actions related to Pegasus.
	Manager struct {
		topics map[string]*topic.Topic[*message.Message]
	}
)

// NewManager returns a new Manager
func NewManager() *Manager {
	return &Manager{
		topics: map[string]*topic.Topic[*message.Message]{},
	}
}

// GetTopicNames returns a list of names of all created Topics.
func (m *Manager) GetTopicNames() []string {
	topics := []string{}

	for k := range m.topics {
		topics = append(topics, k)
	}

	return topics
}

// GetTopicNames returns a list of names of all created Topics.
func (m *Manager) GetTopic(topicName string) (*topic.Topic[*message.Message], error) {
	t, ok := m.topics[topicName]
	if !ok {
		return nil, fmt.Errorf(errTopicNotFound, topicName)
	}

	return t, nil
}

const (
	errTopicAlreadyExists = "topic named '%s' already exists"
	errTopicNotFound      = "topic named '%s' not found"
)

// NewTopic creates a new Topic and adds it to the Manager.
func (m *Manager) NewTopic(name string) error {
	if _, ok := m.topics[name]; ok {
		return fmt.Errorf(errTopicAlreadyExists, name)
	}
	createdTopic, err := topic.NewTopic[*message.Message](name)
	if err != nil {
		return err
	}

	m.topics[createdTopic.Name] = createdTopic
	return nil
}

// Send inserts a message into Topic's internal queue.
func (m *Manager) Send(msg *message.Message) error {
	if err := msg.Validate(); err != nil {
		return err
	}
	topic, ok := m.topics[msg.TopicName]
	if !ok {
		return fmt.Errorf(errTopicNotFound, msg.TopicName)
	}

	return topic.Send(msg)
}

// Pop retrieves a message from Topic's internal queue.
func (m *Manager) Pop(name string) (*message.Message, error) {
	topic, ok := m.topics[name]
	if !ok {
		return nil, fmt.Errorf(errTopicNotFound, name)
	}

	return topic.Pop()
}
