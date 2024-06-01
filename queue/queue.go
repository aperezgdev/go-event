package queue

import (
	"github.com/aperezgdev/go-event/event"
)

type Queue interface {
	Enqueue(...event.Event)
	Dequeue() event.Event
}

type MemoryQueue struct {
	events []event.Event
	Size   uint
}

func InitMemoryQueue() MemoryQueue {
	return MemoryQueue{
		nil,
		0,
	}
}

func (qe *MemoryQueue) Enqueue(ev ...event.Event) {
	if qe.events == nil {
		qe.events = ev
		return
	}

	qe.events = append(qe.events, ev...)
}

func (qe *MemoryQueue) Dequeue() event.Event {
	if len(qe.events) > 0 {
		firstEvent := qe.events[0]
		qe.events = qe.events[1:]
		return firstEvent
	}
	return nil
}
