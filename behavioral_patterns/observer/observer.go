package main

import "fmt"

/**
DIFFERENCE BETWEEN OBSERVER AND PUB-SUB PATTERNS:

- In the OBSERVER pattern, the observers are aware of the Subject .
The Subject maintains a record of the Observers.

- Whereas, in PUBLISHER-SUBSCRIBER, publishers and subscribers don't need to know each other.
They simply communicate with the help of message queues or a broker.
 */

/*
SUBJECT-PUBLISHER INTERFACE: Stream Source of Events
 */
type Subject interface {
	attach(observer Observer)//subscribe
	detach(observer Observer)//unsubscribe
	notifyAll()
}
/*
CONCRETE SUBJECT-PUBLISHER
*/
type subject struct {
	observerList map[int]Observer
	state        string
}
func (s *subject) attach(o Observer) {
	s.observerList[o.Id()] = o
}
func (s *subject) detach(o Observer) {
	delete(s.observerList, o.Id())
}
func (s *subject) notifyAll() {
	for _, observer := range s.observerList {
		observer.update(s.state)
	}
}
func (s *subject) setState(state string) {
	s.state = state
	s.notifyAll()
}

/*
OBSERVER-SUBSCRIBER INTERFACE: Sink of Events Listeners
 */
type Observer interface {
	update(string)
	Id() int
}
type observer struct {
	id int
}
func (o *observer) Id() int {
	return o.id
}
/*
CONCRETE OBSERVERS-SUBSCRIBERS
 */
type observerOne struct {observer}
func (oo *observerOne) update(subjectState string) {
	fmt.Printf("observerOne with ID %d has received new state: %s\n", oo.id, subjectState)
}
type observerTwo struct {observer}
func (ot *observerTwo) update(subjectState string) {
	fmt.Printf("observerTwo with ID %d has received new state: %s\n", ot.id, subjectState)
}


/*
CLIENT
 */
func main() {
	subject := &subject{
		make(map[int]Observer),
		"Initial State",
	}
	observerOne := &observerOne{observer{ id: 1 }}
	observerTwo := &observerTwo{observer{ id: 2 }}

	subject.attach(observerOne)
	subject.detach(observerOne)

	subject.attach(observerTwo)

	subject.setState("New State")
}