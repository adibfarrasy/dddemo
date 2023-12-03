package ddd

type Aggregate interface {
	Entity
	AddEvent(event Event)
	Events() []Event
}

type AggregateBase struct {
	ID     string
	events []Event
}

func (a *AggregateBase) AddEvent(event Event) {
	a.events = append(a.events, event)
}

func (a AggregateBase) Events() []Event {
	return a.events
}
