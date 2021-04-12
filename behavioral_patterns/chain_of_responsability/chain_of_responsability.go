package main

import "fmt"

/*
REQUEST
 */
type request struct {
	name             string
	handlerOneDone   bool
	handlerTwoDone   bool
	handlerThreeDone bool
}

/*
HANDLER INTERFACE & BASE HANDLER
 */
type Handler interface {
	execute(*request)
	setNext(Handler)
}
type baseHandler struct {
	next Handler
}
func (bh *baseHandler) setNext(next Handler) {
	bh.next = next
}

/*
CONCRETE HANDLERS 1, 2, 3
*/
type handlerOne struct {
	baseHandler
}
func (ho *handlerOne) execute(r *request) {
	if r.handlerOneDone {
		fmt.Println("handlerOne already done")
		ho.next.execute(r)
		return
	}
	fmt.Println("doing handlerOne")
	r.handlerOneDone = true
	ho.next.execute(r)
}

type handlerTwo struct {
	baseHandler
}
func (ht *handlerTwo) execute(r *request) {
	if r.handlerTwoDone {
		fmt.Println("handlerTwo already done")
		ht.next.execute(r)
		return
	}
	fmt.Println("doing handlerTwo")
	r.handlerTwoDone = true
	ht.next.execute(r)
}

type handlerThree struct {
	baseHandler
}
func (hth *handlerThree) execute(r *request) {
	if r.handlerThreeDone {
		fmt.Println("handlerThree already done")
		return
	}
	fmt.Println("doing handlerThree")
	r.handlerThreeDone = true
}


/*
CLIENT
 */
func main() {
	//SETUP CHAIN OF RESPONSIBILITY
	handlerThree := new(handlerThree)
	handlerTwo := new(handlerTwo)
	handlerTwo.setNext(handlerThree)
	handlerOne := new(handlerOne)
	handlerOne.setNext(handlerTwo)

	//CREATE REQUEST
	request := &request{name: "abc"}

	//EXECUTE REQUEST IN CHAIN
	handlerOne.execute(request)
}

