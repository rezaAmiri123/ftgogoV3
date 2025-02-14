package am

// import (
// 	"context"
// 	"strings"

// 	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
// 	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
// 	"google.golang.org/protobuf/proto"
// )

// type (
// 	CommandMessageHandler     = MessageHandler[IncomingCommandMessage]
// 	CommandMessageHandlerFunc func(ctx context.Context, msg IncomingCommandMessage) error

// 	Command interface {
// 		ddd.Command
// 		Destination() string
// 	}
// 	command struct {
// 		ddd.Command
// 		destination string
// 	}
// )

// func NewCommand(name, destination string, payload ddd.CommandPayload, options ...ddd.CommandOption) Command {
// 	return command{
// 		Command:     ddd.NewCommand(name, payload, options...),
// 		destination: destination,
// 	}
// }

// func (c command) Destination() string {
// 	return c.destination
// }

// func (f CommandMessageHandlerFunc) HandleMessage(ctx context.Context, cmd IncomingCommandMessage) error {
// 	return f(ctx, cmd)
// }


// func NewCommandMessageHandler(reg registry.Registry, publisher ReplyPublisher, handler ddd.CommandHandler[ddd.Command]) RawMessageHandler {
// 	return commandMsgHandler{
// 		reg:       reg,
// 		publisher: publisher,
// 		handler:   handler,
// 	}
// }

// func (h commandMsgHandler) HandleMessage(ctx context.Context, msg IncomingRawMessage) error {
// 	var commandData CommandMessageData

// 	err := proto.Unmarshal(msg.Data(), &commandData)
// 	if err != nil {
// 		return err
// 	}

// 	commandName := msg.MessageName()

// 	payload, err := h.reg.Deserialize(commandName, commandData.GetPayload())
// 	if err != nil {
// 		return err
// 	}

// 	commandMsg := commandMessage{
// 		id:         msg.ID(),
// 		name:       commandName,
// 		payload:    payload,
// 		metadata:   commandData.Metadata.AsMap(),
// 		occurredAt: commandData.OccurredAt.AsTime(),
// 		msg:        msg,
// 	}

// 	destination := commandMsg.Metadata().Get(CommandReplyChannelHdr).(string)

// 	reply, err := h.handler.HandleCommand(ctx, commandMsg)
// 	if err != nil {
// 		return h.publishReply(ctx, destination, h.failure(reply, commandMsg))
// 	}
// 	return h.publishReply(ctx, destination, h.success(reply, commandMsg))
// }

