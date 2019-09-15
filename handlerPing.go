package eventdemo

import (
	"fmt"
	"sync"
	"time"
)

// PingHandler ...
type PingHandler struct{}

// HandleMessage ...
func (handler PingHandler) HandleMessage(msg Message, wg *sync.WaitGroup) {
	time.Sleep(250 * time.Millisecond) // simulate hard word
	fmt.Printf("PING   | %s\n", msg)
	wg.Done()
}
