package demo

import (
	"sync"
	"time"
)

// Generator ...
type Generator struct {
	waitGroup *sync.WaitGroup
}

// NewGenerator ...
func NewGenerator(waitGroup *sync.WaitGroup) Generator {
	return Generator{waitGroup}
}

// GenerateMessages ...
func (g *Generator) GenerateMessages(messages chan Message) {
	defer close(messages)

	for _, message := range produceMessages() {
		messages <- message
		time.Sleep(10 * time.Millisecond) // simulate hard work
	}

	g.waitGroup.Done()
}

func produceMessages() []Message {
	messages := []Message{}

	messages = append(messages, NewMessage("Hello Gophers", Hello))
	messages = append(messages, NewMessage("Hello Stuttgart", Hello))

	messages = append(messages, NewMessage("First steps in go", Ping))
	messages = append(messages, NewMessage("Next steps in go", Ping))
	messages = append(messages, NewMessage("I am still listening", Ping))
	messages = append(messages, NewMessage(`{ "content-type": "json" }`, Ping))
	messages = append(messages, NewMessage("Can't wait to try it out myself", Ping))
	messages = append(messages, NewMessage("Super excited to hear the next talk", Ping))

	messages = append(messages, NewMessage("Thanks for coming", Goodbye))
	messages = append(messages, NewMessage("Good night Stuttgart", Goodbye))

	for i := 0; i < 100; i++ {
		messages = append(messages, messages[i%10])
	}
	return messages
}
