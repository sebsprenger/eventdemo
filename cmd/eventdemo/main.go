package main

import (
	"sync"

	demo "github.com/sebsprenger/eventdemo"
)

func main() {
	var waitGroup sync.WaitGroup

	generator := demo.NewGenerator(&waitGroup)
	filter := demo.NewFilter(&waitGroup)
	dispatcher := demo.NewDispatcher(&waitGroup)

	dispatcher.RegisterHandler(demo.Hello, demo.HelloHandler{})
	dispatcher.RegisterHandler(demo.Ping, demo.PingHandler{})
	dispatcher.RegisterHandler(demo.Goodbye, demo.GoodByeHandler{})

	messageChannel := make(chan demo.Message, 200)
	filteredChannel := make(chan demo.Message, 200)

	waitGroup.Add(3)

	generator.GenerateMessages(messageChannel)
	filter.Filter(messageChannel, filteredChannel, demo.Hello, demo.Ping, demo.Goodbye)
	dispatcher.Dispatch(filteredChannel)

	waitGroup.Wait()
}
