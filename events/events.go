package events

import (
	"io"
	"os"
	"sync"
	"time"
)

type EventType string

type Event interface{}

type Handler func(args Event)

var eventMap = make(map[EventType][]Handler)

var eventsChan chan struct {
	Event
	EventType
}

var rm sync.RWMutex

func Subscribe(h Handler, e EventType) {

	if eventMap[e] == nil {
		eventMap[e] = make([]Handler, 0)
	}

	eventMap[e] = append(eventMap[e], h)

}

func notify(t EventType, e Event) {

	for _, h := range eventMap[t] {
		go h(e)
	}
}

func PublishEvent(t EventType, e Event) {
	rm.RLock()

	go func() {

		eventsChan <- struct {
			Event
			EventType
		}{Event: e, EventType: t}

	}()

	rm.RUnlock()
}

// to be run as a goroutine
func Run() {
	eventsChan = make(chan struct {
		Event
		EventType
	})

	defer close(eventsChan)

	for {

		select {
		case evt := <-eventsChan:

			notify(evt.EventType, evt.Event)
		}

		time.Sleep(time.Millisecond)
	}
}

func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
