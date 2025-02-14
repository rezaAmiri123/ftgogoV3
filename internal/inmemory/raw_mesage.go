package inmemory

import (
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/am"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type rawMessage struct {
	msg am.Message
}

var _ am.Message = (*rawMessage)(nil)

func (m rawMessage) ID() string              { return m.msg.ID() }
func (m rawMessage) Subject() string         { return m.msg.Subject() }
func (m rawMessage) MessageName() string     { return m.msg.MessageName() }
func (m rawMessage) Data() []byte            { return m.msg.Data() }
func (m *rawMessage) Metadata() ddd.Metadata { return m.msg.Metadata() }
func (m *rawMessage) SentAt() time.Time      { return m.msg.SentAt() }
func (m *rawMessage) ReceivedAt() time.Time  { return time.Now() }
func (m *rawMessage) Ack() error             { return nil }
func (m *rawMessage) NAck() error            { return nil }
func (m rawMessage) Extend() error           { return nil }
func (m *rawMessage) Kill() error            { return nil }
