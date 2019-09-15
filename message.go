package demo

import (
	"fmt"
	"time"
)

// MessageType identifies available types of messages
type MessageType int

// Available MessageTypes
const (
	Hello   MessageType = iota
	Ping    MessageType = iota
	Goodbye MessageType = iota
)

// Message is a container for an arbitrary message
type Message struct {
	Content      string
	CreationTime time.Time
	Type         MessageType
}

// NewMessage create a new message and sets the creation time to "now"
func NewMessage(content string, msgType MessageType) Message {
	return Message{
		Content:      content,
		CreationTime: time.Now(),
		Type:         msgType,
	}
}

// GetAge calculates the time passed since message creation
func (m Message) GetAge() time.Duration {
	return time.Now().Sub(m.CreationTime)
}

func (m Message) String() string {
	return fmt.Sprintf("%d | %-15s | %s", m.Type, m.GetAge(), m.Content)
}
