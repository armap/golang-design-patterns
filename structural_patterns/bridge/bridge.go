package main

import "fmt"

/*
ABSTRACTION INTERFACE
 */
type Abstraction interface {
	methodA()
}
type abstraction struct {
	implementation Implementation
}

/*
CONCRETE ABSTRACTIONS One & Two
 */
func NewConcreteAbsOne(i Implementation) concreteAbsOne {
	return concreteAbsOne{ abstraction{ i } }
}
type concreteAbsOne struct {
	abstraction
}
func (m *concreteAbsOne) methodA() {
	fmt.Println("Abstraction ONE:")
	m.implementation.implementMethodA()
}
func NewConcreteAbsTwo(i Implementation) concreteAbsTwo {
	return concreteAbsTwo{ abstraction{ i } }
}
type concreteAbsTwo struct {
	abstraction
}
func (w *concreteAbsTwo) methodA() {
	fmt.Println("Abstraction TWO:")
	w.implementation.implementMethodA()
}

/*
IMPLEMENTATION INTERFACE
*/
type Implementation interface {
	implementMethodA()
}

/*
CONCRETE IMPLEMENTATIONS Blue & Red
*/
type concreteImplBlue struct {
}
func (i *concreteImplBlue) implementMethodA() {
	fmt.Println("BLUE implementation for method A")
}
type concreteImplRed struct {
}
func (i *concreteImplRed) implementMethodA() {
	fmt.Println("RED implementation for method A")
}


/*
CLIENT
*/
func main() {
	oneBlue := NewConcreteAbsOne(new(concreteImplBlue))
	oneBlue.methodA()
	oneRed := NewConcreteAbsOne(new(concreteImplRed))
	oneRed.methodA()

	twoBlue := NewConcreteAbsTwo(new(concreteImplBlue))
	twoBlue.methodA()
	twoRed := NewConcreteAbsTwo(new(concreteImplRed))
	twoRed.methodA()
}