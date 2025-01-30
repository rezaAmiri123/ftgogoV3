package ddd

const (
	AggregateNameKey    = "aggregate-name"
	AggregateIDKey      = "aggregate-id"
	AggregateVersionKey = "aggregate-vesrion"
)

type (
	AggregateNamer interface {
		AggregateName() string
	}

	Eventer interface {
		AddEvent(string, EventPayload, ...EventOption)
		Events() []AggregateEvent
		ClearEvents()
	}

	aggregate struct {
		Entity
		events []AggregateEvent
	}

	AggregateEvent interface {
		Event
		AggregateName() string
		AggregateID() string
		AggregateVersion() int
	}

	aggregateEvent struct {
		event
	}

	Aggregate interface{
		IDer
		AggregateNamer
		Eventer
		IDSetter
		NameSetter
	}
)

var _ Aggregate =(*aggregate)(nil)

func NewAggregate(id, name string) *aggregate {
	return &aggregate{
		Entity: NewEntity(id, name),
		events: make([]AggregateEvent, 0),
	}
}

func (a aggregate) AggregateName() string    { return a.EntityName() }
func (a aggregate) Events() []AggregateEvent { return a.events }
func (a aggregate) ClearEvents()             { a.events = []AggregateEvent{} }

func (a *aggregate) AddEvent(name string, payload EventPayload, options ...EventOption) {
	options = append(options, Metadata{
		AggregateNameKey: a.EntityName(),
		AggregateIDKey:   a.ID(),
	})

	a.events = append(a.events, aggregateEvent{
		event: newEvent(name, payload, options...),
	})
}

func (a *aggregate) setEvents(events []AggregateEvent) { a.events = events }

func (a aggregateEvent) AggregateName() string { return a.metadata.Get(AggregateNameKey).(string) }
func (a aggregateEvent) AggregateID() string   { return a.metadata.Get(AggregateIDKey).(string) }
func (a aggregateEvent) AggregateVersion() int { return a.metadata.Get(AggregateVersionKey).(int) }
