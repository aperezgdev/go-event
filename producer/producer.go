package producer

import "github.com/aperezgdev/go-event/event"

type Producer func(...interface{}) event.Event
