package am

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type (
	Message interface {
		ddd.IDer
		MessageName() string
		Ack() error
		NAck() error
		Extend() error
		Kill() error
	}

	MessageHandler[O Message] interface {
		HandleMessage(ctx context.Context, msg O) error
	}

	MessageHandlerFunc[O Message] func(ctx context.Context, msg O) error

	MessagePulisher[I any] interface {
		Publish(ctx context.Context, topicName string, v I) error
	}

	MessageSubscrier[O Message] interface {
		Subscribe(topicName string, handler MessageHandler[O], options ...SubscriberOption) error
	}

	MessageStream[I any, O Message] interface {
		MessagePulisher[I]
		MessageSubscrier[O]
	}
)

func (f MessageHandlerFunc[O]) HandleMessage(ctx context.Context, msg O) error {
	return f(ctx, msg)
}
