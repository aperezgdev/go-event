package examples

import (
	"errors"
	"fmt"

	"github.com/aperezgdev/go-event/event"
)

func UserCreatedConsumer(e event.Event) error {
	ev, ok := e.(UserCreated)

	if !ok {
		return errors.New("could not cast event to user created")
	}

	fmt.Println("Sending email to " + ev.userEmail)

	return nil
}
