package examples

import (
	"errors"

	"github.com/aperezgdev/go-event/event"
)

func ErrorConsumer(ev event.Event) error {
	return errors.New("error has ocurred")
}
