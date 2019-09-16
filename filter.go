package eventdemo

import (
	"sync"
)

// Filter ...
type Filter struct {
	waitGroup *sync.WaitGroup
}

// NewFilter ...
func NewFilter(wg *sync.WaitGroup) Filter {
	return Filter{wg}
}

// Filter ...
func (f Filter) Filter(input chan Message, output chan Message, msgTypes ...MessageType) {
	defer close(output)

	for message := range input {
		for _, msgType := range msgTypes {
			if message.Type == msgType {
				output <- message
				continue
			}
		}
	}

	f.waitGroup.Done()
}
