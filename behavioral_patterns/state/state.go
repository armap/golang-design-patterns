package main

import "fmt"

/*
CONTEXT
 */
type contextFan struct {
	stateHigh state
	stateLow  state
	stateOff  state

	state state
}
func NewContextFan() *contextFan {
	cf := new(contextFan)
	cf.stateHigh = &stateHigh{cf}
	cf.stateLow = &stateLow{cf}
	cf.stateOff = &stateOff{cf}

	//Set the Fan's initial state to OFF
	cf.setState(cf.stateOff)

	return cf
}
func (cf *contextFan) setState(s state) {
	cf.state = s
}
func (cf *contextFan) request() {
	cf.state.handleRequest()
}

/*
STATE INTERFACE
 */
type state interface {
	handleRequest()
}
/*
CONCRETE STATES (OFF, LOW, HIGH)
*/
type stateOff struct {
	contextFan *contextFan
}
func (so *stateOff) handleRequest() {
	fmt.Println("Fan state is OFF, so turning it on with low speed. Setting state to LOW")
	so.contextFan.setState(so.contextFan.stateLow)
}
type stateLow struct {
	contextFan *contextFan
}
func (sl *stateLow) handleRequest() {
	fmt.Println("Fan state is LOW, so increasing speed to high. Setting state to HIGH")
	sl.contextFan.setState(sl.contextFan.stateHigh)
}
type stateHigh struct {
	contextFan *contextFan
}
func (sh *stateHigh) handleRequest() {
	fmt.Println("Fan state is HIGH, so turning it down. Setting state to OFF")
	sh.contextFan.setState(sh.contextFan.stateOff)
}


/*
CLIENT
 */
func main() {
	//Initialize Fan switched Off
	fan := NewContextFan()

	//Turning on Fan at Low speed
	fan.request()

	//Increase Fan speed to High
	fan.request()

	//Turn Off Fan
	fan.request()
}