package observer

import "fmt"

// Subject declare a subject struct
type Subject struct {
	observers []Observer
	context   string
}

// NewSubject serve caller to create a Subject
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// Subscribe serve caller to subscribe a Subject
func (s *Subject) Subscribe(o Observer) {
	s.observers = append(s.observers, o)
}

// UpdateContext serve caller to update subject's context
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// Reader declare a reader struct
type Reader struct {
	name string
}

// NewReader serve caller to create a Reader
func NewReader(name string) *Reader {
	return &Reader{name: name}
}

// Update serve caller to received message from Subject
func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

// Observer declare a observer interface
type Observer interface {
	Update(*Subject)
}
