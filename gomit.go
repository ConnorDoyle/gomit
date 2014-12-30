package gomit

import (
	// "fmt"
	"time"
)

type Emitter struct {
	Handlers []func(Event)
}

type SourceEvent interface {
}

type Event struct {
	Header      EventHeader
	SourceEvent SourceEvent
}

type EventHeader struct {
	Time     time.Time
	Metadata map[string]string
}

func NewEvent(s SourceEvent, m map[string]string) Event {
	if m == nil {
		m = make(map[string]string)
	}
	h := EventHeader{time.Now(), m}

	event := Event{h, s}
	return event
}

func (e *Emitter) Fire(event Event) {
	for _, f := range e.Handlers {
		go f(event)
	}
}

func (e *Emitter) AddHandler(f func(Event)) {
	e.Handlers = append(e.Handlers, f)
}
