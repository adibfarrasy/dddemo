package ddd

type Entity interface {
	ID() string
}

type EntityBase struct {
	id string
}

func (e EntityBase) ID() string {
	return e.id
}
