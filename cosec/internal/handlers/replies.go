package handlers

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/cosec/internal"
	"github.com/rezaAmiri123/ftgogoV3/internal/am"
)

func RegisterReplyHandlers(subscriber am.RawMessageStream, replyHandlers am.RawMessageHandler) (err error) {
	replyMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) error {
		return replyHandlers.HandleMessage(ctx, msg)
	})

	_, err = subscriber.Subscribe(internal.CreateOrderReplyChannel, replyMsgHandler, am.GroupName("cosec-replies"))
	return err
}
