package main

import (
	"errors"
	"log"
)

// event types the client want to be notified
const (
	Available  = iota
	OutOfStock
)

var (
	ErrNoListeners = errors.New("No listeners for the given type")
	ErrTypeNotRegistered = errors.New("Type not registered")
)

type EventManager struct {
	listenersByType map[uint][]Listener // these are the subscribers
}

func NewEventManager() *EventManager {
	em := &EventManager{}
	em.listenersByType = make(map[uint][]Listener)
	return em
}

func (em *EventManager) AddListener(listener Listener, eventType uint) {
	listeners, ok := em.listenersByType[eventType]
	if !ok {
		listeners = make([]Listener, 1)
		listeners[0] = listener
	} else {
		listeners = append(listeners, listener)
	}

	em.listenersByType[eventType] = listeners
}

func (em *EventManager) RemoveListener(listener Listener, eventType uint) error {
	listeners, ok := em.listenersByType[eventType]
	if !ok {
		return ErrTypeNotRegistered
	}

	if len(listeners) == 0 {
		return ErrNoListeners
	}

	listeners = em.filterListenerOut(listeners, listener)
	em.listenersByType[eventType] = listeners

	return nil
}

func (em *EventManager) filterListenerOut(listeners []Listener, listener Listener) []Listener {
	nListeners := make([]Listener, 0, len(listeners)-1)
	for _, l := range listeners {
		if l != listener {
			nListeners = append(nListeners, l)
		}
	}
	return nListeners
}

func (em *EventManager) Notify(eventType uint) {
	listeners, ok := em.listenersByType[eventType]
	if ok {
		for _, listener := range listeners {
			listener.Notify()
		}
	}
}

type Listener interface {
	Notify()
}

type EmailListener struct {
	Email string
}

func (e EmailListener) Notify() {
	log.Printf("sending email to %v...", e.Email)
}

func main() {
	em := NewEventManager()
	emailListener1 := EmailListener{Email: "danilo@gmail.com"}
	emailListener2 := EmailListener{Email: "fitz@gmail.com"}
	emailListener3 := EmailListener{Email: "nina@gmail.com"}
	emailListener4 := EmailListener{Email: "aurora@gmail.com"}

	em.AddListener(emailListener1, Available)
	em.AddListener(emailListener2, Available)
	em.AddListener(emailListener3, Available)
	em.AddListener(emailListener4, OutOfStock)

	log.Printf("%#v\n", em.listenersByType)
	if err := em.RemoveListener(emailListener3, OutOfStock); err != nil {
		log.Printf("Error removing listener %v\n", err)
	}

	log.Printf("%#v\n", em.listenersByType)

	em.Notify(Available)
}
