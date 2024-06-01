package handler

import (
	"fmt"

	"github.com/aperezgdev/go-event/dispatcher"
	"github.com/aperezgdev/go-event/queue"
)

type EventsHandler struct {
	Queue      queue.Queue
	Dispatcher dispatcher.Dispatcher
}

func (eh EventsHandler) HandlerEvents() {
	for {
		e := eh.Queue.Dequeue()

		if e == nil {
			break
		}

		err := eh.Dispatcher.Dispatch(e)

		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
