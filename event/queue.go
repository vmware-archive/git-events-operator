package event

import (
	"fmt"
	"sync"
)

type Queue struct {
	events  []Event
	brokers []Broker
}

type EventImplementations []Event

func NewQueue(brokers []Broker) *Queue {
	return &Queue{
		brokers: brokers,
	}

}

func (q *Queue) ConcurrentStart() chan error {
	errch := make(chan error)
	go func() {
		for _, broker := range q.brokers {
			errch2 := <-broker.ConcurrentWatch(q)

			// Concurrently pass each error to the broader error channel
			// as we see one, otherwise block on the channel.
			go func() {
				for {
					errch <- errch2
				}
			}()
		}
	}()
	return errch
}

var (
	addMutex sync.Mutex
	popMutex sync.Mutex
)

// AddEvent is a thread safe way to add an event to the queue.
// Use this function to add an event to the queue, it will processed FIFO
func (q *Queue) AddEvent(event Event) {
	addMutex.Lock()
	defer addMutex.Unlock()

	// Append to the end of the queue
	q.events = append(q.events, event)
}

// PopFromQueue is a thread safe way to pop the oldest event from the queue
// The queue behaves as a FIFO queue, and right now that cannot be configured.
// Use this function to take the oldest message out of the queue to operate on it.
// PopFromQueue will return an error if there is nothing in the queue.
func (q *Queue) PopFromQueue() (Event, error) {
	popMutex.Lock()
	defer popMutex.Unlock()
	var e Event
	if len(q.events) == 0 {
		return e, fmt.Errorf("empty queue")
	}
	e, q.events = q.events[0], q.events[1:]
	return e, nil
}

// TODO @kris-nova
// Right now the queue is relatively "dumb" for lack of a better term.
// It might make sense to have the operator portion of the code "ACK"
// an event after it's been popped off the queue.
// For now instead of going with the "ACK" approach, we can just have
// the operator re-add it to queue and hope for it to be processed later.
// Actually that might just flood the queue. We should probably just ignore
// it. Whatever. We need to figure this out.
