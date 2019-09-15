package demo

import (
	"fmt"
	"sync"
	"time"
)

// GoodByeHandler ...
type GoodByeHandler struct{}

// HandleMessage ...
func (handler GoodByeHandler) HandleMessage(msg Message, wg *sync.WaitGroup) {
	time.Sleep(40 * time.Millisecond) // simulate hard word
	fmt.Printf("BYEBYE | %s\n", msg)
	wg.Done()
}
