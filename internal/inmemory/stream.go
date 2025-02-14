package inmemory

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
)

type stream struct {
	subscriptions map[string][]am.MessageHandler
}

var _ am.MessageStream = (*stream)(nil)

func NewStream() stream {
	return stream{
		subscriptions: make(map[string][]am.MessageHandler),
	}
}

func (t stream) Publish(ctx context.Context, topicName string, v am.Message) error {
	for _, handler := range t.subscriptions[topicName] {
		err := handler.HandleMessage(ctx, &rawMessage{v})
		if err != nil {
			return err
		}
	}
	return nil
}

func (t stream) Subscribe(topicName string, handler am.MessageHandler, options ...am.SubscriberOption) (am.Subscription, error) {
	cfg := am.NewSubscriberConfig(options)

	var filters map[string]struct{}
	if len(cfg.MessageFilters()) > 0 {
		filters = make(map[string]struct{})
		for _, key := range cfg.MessageFilters() {
			filters[key] = struct{}{}
		}
	}

	fn := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) error {
		if filters != nil {
			if _, exists := filters[msg.MessageName()]; !exists {
				return nil
			}
		}

		return handler.HandleMessage(ctx, msg)
	})

	t.subscriptions[topicName] = append(t.subscriptions[topicName], fn)

	return nil, nil
}

func (t stream) Unsubscribe() error {
	return nil
}
