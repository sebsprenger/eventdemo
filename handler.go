package eventdemo

import (
	"sync"
)

// MessageHandler ...
type MessageHandler interface {
	HandleMessage(Message, *sync.WaitGroup)
}
