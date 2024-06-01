package main

import (
	"fmt"

	"github.com/aperezgdev/go-event/dispatcher"
	"github.com/aperezgdev/go-event/examples"
	"github.com/aperezgdev/go-event/queue"
)

func registerUser(email string) examples.UserCreated {
	fmt.Println("Registrando usuario")

	return examples.InitUserCreated(email)
}

func main() {
	var event = registerUser("example@example.com")

	var event2 = registerUser("example2@example2.com")

	dispatch := dispatcher.InitDefaultEventDispatcher()

	dispatch.On(examples.USER_CREATED, examples.UserCreatedConsumer, examples.ErrorConsumer)

	q := queue.InitQueue()

	q.Enqueue(event)

	q.Enqueue(event2)

	for {
		ev := q.Dequeue()

		if ev == nil {
			break
		}

		err := dispatch.Dispatch(ev)

		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
