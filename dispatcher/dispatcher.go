package dispatcher

import (
	"errors"
	"fmt"

	"github.com/aperezgdev/go-event/consumer"
	"github.com/aperezgdev/go-event/event"
)

type Dispatcher interface {
	Dispatch(e event.Event)
	On(eventName string, toAddConsumer consumer.Consumer)
	Remove(eventName string)
}

type DefaultEventDispatcher struct {
	relations map[string][]consumer.Consumer
}

func InitDefaultEventDispatcher() DefaultEventDispatcher {
	return DefaultEventDispatcher{
		relations: make(map[string][]consumer.Consumer),
	}
}

func (ed *DefaultEventDispatcher) Dispatch(e event.Event) error {
	if ed.relations[e.Name()] == nil {
		return errors.New("not exists comsumer for event " + e.Name())
	}

	for _, consumer := range ed.relations[e.Name()] {
		err := consumer(e)

		if err != nil {
			fmt.Println("error has ocurred at the event " + e.Name())
		}
	}

	return nil
}

func (ed *DefaultEventDispatcher) On(eventName string, toAddConsumer consumer.Consumer) {
	if ed.relations[eventName] == nil {
		ed.relations[eventName] = []consumer.Consumer{toAddConsumer}
		return
	}

	ed.relations[eventName] = append(ed.relations[eventName], toAddConsumer)
}

func (ed *DefaultEventDispatcher) Remove(eventName string) {
	delete(ed.relations, eventName)
}
