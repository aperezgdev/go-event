package main

import (
	"github.com/aperezgdev/go-event/dispatcher"
	"github.com/aperezgdev/go-event/examples"
	"github.com/aperezgdev/go-event/handler"
	"github.com/aperezgdev/go-event/queue"
)

func main() {
	ev := examples.InitUserCreated("example@example.com")

	ev2 := examples.InitUserCreated("example2@example2.com")

	dispatch := dispatcher.InitDefaultEventDispatcher()

	dispatch.On(examples.USER_CREATED, examples.UserCreatedConsumer, examples.ErrorConsumer)

	q := queue.InitMemoryQueue()

	h := handler.EventsHandler{
		Queue:      &q,
		Dispatcher: &dispatch,
	}

	q.Enqueue(ev, ev2)

	h.HandlerEvents()
}
