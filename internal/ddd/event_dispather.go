package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface{
	Subscribe(event Event, handler EventHandler)
}

//go:generate mockery --name EventPublisher
type EventPublisher interface{
	Publish(ctx context.Context, events ...Event)error
}

type EventDispatcher struct{
	handlers map[string][]EventHandler
	mu sync.Mutex
}

var _ interface{
	EventSubscriber
	EventPublisher
} = (*EventDispatcher)(nil)

func NewEventDispatcher()*EventDispatcher{
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher)Subscribe(event Event, handler EventHandler){
	d.mu.Lock()
	defer d.mu.Unlock()

	d.handlers[event.EventName()] = append(d.handlers[event.EventName()], handler)
}

func (d *EventDispatcher)Publish(ctx context.Context, events ...Event)(err error){
	for _, event := range events{
		for _, handler := range d.handlers[event.EventName()]{
			err = handler(ctx, event)
			if err!= nil{
				return err
			} 
		}
	}
	return nil
}
