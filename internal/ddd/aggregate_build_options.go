package ddd

type EventSetter interface{
	setEvents([]Event)
}