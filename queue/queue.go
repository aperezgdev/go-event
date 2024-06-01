package queue

import (
	"github.com/aperezgdev/go-event/event"
)

type Queue struct {
	events []event.Event
	Size   uint
}

func InitQueue() Queue {
	return Queue{
		nil,
		0,
	}
}

func (qe *Queue) Enqueue(ev event.Event) {
	if qe.events == nil {
		qe.events = []event.Event{ev}
		return
	}

	qe.events = append(qe.events, ev)
}

func (qe *Queue) Dequeue() event.Event {
	if len(qe.events) > 0 {
		firstEvent := qe.events[0]
		qe.events = qe.events[1:]
		return firstEvent
	}
	return nil
}
