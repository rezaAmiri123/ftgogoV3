package tm

import (
	"context"
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/stackus/errors"
)

type ErrDuplicateMessage string

func (e ErrDuplicateMessage) Error() string {
	return fmt.Sprintf("dublicate message id encountered: %s", string(e))
}

type InboxStore interface {
	Save(ctx context.Context, msg am.IncomingMessage) error
}

func InboxHandler(store InboxStore) am.MessageHandlerMiddleware {
	return func(next am.MessageHandler) am.MessageHandler {
		return am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) error {
			// try to insert the message
			err := store.Save(ctx, msg)
			if err != nil {
				var errDupe ErrDuplicateMessage
				if errors.As(err, &errDupe) {
					// duplicate message; return without an error to let the message Ack
					return nil
				}
				return err
			}
			return next.HandleMessage(ctx, msg)
		})
	}
}
