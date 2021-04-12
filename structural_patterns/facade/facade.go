package main

import (
	"fmt"
	"log"
)

/*
FACADE - simplifies access interface to subsystems
 */
type Facade struct {
	subsystemA *subsystemA
	subsystemB *subsystemB
	subsystemC *subsystemC
	subsystemD *subsystemC
}

func NewFacade(id string) *Facade {
	fmt.Println("Creating Facade")
	facade := &Facade{
		subsystemA: newSubsystemA(id),
		subsystemB: newSubsystemB(),
		subsystemD: new(subsystemC),
	}
	return facade
}
func (w *Facade) simpleMethodOne(id string, amount int) error {
	fmt.Println("simpleMethodOne")
	err := w.subsystemA.methodA(id)
	if err != nil {
		return err
	}
	w.subsystemB.methodB1(amount)
	w.subsystemD.methodC1()
	return nil
}
func (w *Facade) simpleMethodTwo(id string, amount int) error {
	fmt.Println("simpleMethodTwo")
	err := w.subsystemA.methodA(id)
	if err != nil {
		return err
	}
	err = w.subsystemB.methodB2(amount)
	if err != nil {
		return err
	}
	w.subsystemD.methodC2()
	return nil
}

/*
COMPLEX SUBSYSTEMS
*/

// SUBSYSTEM A - checks ID
type subsystemA struct {
	ID string
}
func newSubsystemA(id string) *subsystemA {
	return &subsystemA{
		ID: id,
	}
}
func (a *subsystemA) methodA(id string) error {
	if a.ID != id {
		return fmt.Errorf("Id is incorrect")
	}
	fmt.Println("methodA - Id is correct")
	return nil
}

// SUBSYSTEM B - adds amount into total
type subsystemB struct {
	total int
}
func newSubsystemB() *subsystemB {
	return &subsystemB{
		total: 0,
	}
}
func (w *subsystemB) methodB1(amount int) {
	w.total += amount
	fmt.Printf("methodB1 - added %d into total: %d\n", amount, w.total)
	return
}
func (w *subsystemB) methodB2(amount int) error {
	w.total += amount * 2
	fmt.Printf("methodB1 - added double of %d into total: %d\n", amount, w.total)
	w.total = w.total - amount
	return nil
}

// SUBSYSTEM C - prints something
type subsystemC struct {
}
func (n *subsystemC) methodC1() {
	fmt.Println("methodC1")
}
func (n *subsystemC) methodC2() {
	fmt.Println("methodC2")
}


/*
CLIENT
*/
func main() {
	walletFacade := NewFacade("abc")

	err := walletFacade.simpleMethodOne("abc",10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = walletFacade.simpleMethodTwo("abc", 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

