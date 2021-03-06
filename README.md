# eventdemo
An event-driven example to demonstrate the use and impact of goroutines.

## Installation
* `go get github.com/sebsprenger/eventdemo` to download code into `$HOME/go/src/`
* `go run github.com/sebsprenger/eventdemo/cmd/eventdemo` to build and execute

## Overview

There is a generator generating events/messages into a channel.  This channel is read by a filter, filtering all message types by default and passing them in yet another channel. The dispatcher reads the filtered channel and dispatches the messages to different handlers where they are printed. All parts are connected to a `WaitGroup` to synchronize.
Various parts of the code contain a `time.Sleep()` to simulate some work. Adding the `go` keyword in the right places accelerates execution of the program by orders of magnitude.

![Architecture](https://github.com/sebsprenger/eventdemo/blob/master/arch.png)
