package main

import "fmt"

/*
COMPATIBLE THING
*/
type CompatibleThing interface {
	doSomethingCompatible()
}
type compatibleThing struct {}
func (ct *compatibleThing) doSomethingCompatible() {
	fmt.Println("doing something with Compatible Thing")
}

/*
INCOMPATIBLE THING
*/
type incompatibleThing struct{}
func (it *incompatibleThing) doSomethingIncompatible() {
	fmt.Println("..Incompatible Thing")
}

/*
ADAPTER
*/
type adapter struct {
	it *incompatibleThing
}
func (w *adapter) doSomethingCompatible() {
	fmt.Print("doing something with Adapter for..")
	w.it.doSomethingIncompatible()
}


/*
CLIENT
*/
func DoSomething(compatibleThing CompatibleThing) {
	compatibleThing.doSomethingCompatible()
}
func main() {
	compatibleThing := &compatibleThing{}
	DoSomething(compatibleThing)

	incompatibleThing := &incompatibleThing{}
	adapter := &adapter{incompatibleThing }
	DoSomething(adapter)
}