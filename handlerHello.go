package eventdemo

import (
	"fmt"
	"sync"
	"time"
)

// HelloHandler ...
type HelloHandler struct{}

// HandleMessage ...
func (handler HelloHandler) HandleMessage(msg Message, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond) // simulate hard word
	fmt.Printf("HELLO  | %s\n", msg)
	wg.Done()
}
