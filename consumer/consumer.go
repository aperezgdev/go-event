package consumer

import "github.com/aperezgdev/go-event/event"

type Consumer func(e event.Event) error
