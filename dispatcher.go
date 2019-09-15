package eventdemo

import (
	"log"
	"sync"
)

// Dispatcher ...
type Dispatcher struct {
	handlers  map[MessageType]MessageHandler
	waitGroup *sync.WaitGroup
}

// NewDispatcher ...
func NewDispatcher(wg *sync.WaitGroup) Dispatcher {
	handlers := make(map[MessageType]MessageHandler)
	return Dispatcher{handlers, wg}
}

// RegisterHandler ...
func (d Dispatcher) RegisterHandler(msgType MessageType, handler MessageHandler) {
	d.handlers[msgType] = handler
}

// Dispatch ...
func (d Dispatcher) Dispatch(messages chan Message) {
	for message := range messages {
		handler, ok := d.handlers[message.Type]
		if ok != true {
			log.Printf("Did not find a handler for type %d\n", message.Type)
			continue
		}
		d.waitGroup.Add(1)
		handler.HandleMessage(message, d.waitGroup)
	}

	d.waitGroup.Done()
}
