package main

import "fmt"
/**
The TEMPLATE METHOD pattern is an "Inversion Of Control" (IoC) example
 */

/*
ABSTRACT CLASS (INTERFACES AND STRUCT) WITH TEMPLATE METHOD
 */
type IAbstractStep1 interface {
	abstractStep1()
}
type IAbstractStep2 interface {
	abstractStep2()
}
type IHookStep3 interface {
	hookStep3() bool
}
type IHookStep4 interface {
	hookStep4()
}
type IAbstractTemplate interface {
	IAbstractStep1
	IAbstractStep2
	IHookStep3
	IHookStep4
}
type abstractTemplate struct {
	iATemplate IAbstractTemplate
}
func (at *abstractTemplate) templateMethod() {
	at.iATemplate.abstractStep1()
	at.iATemplate.abstractStep2()
	if at.iATemplate.hookStep3() {
		at.iATemplate.hookStep4()
	}
}
func (at *abstractTemplate) hookStep3() bool {
	fmt.Println("Default implementation hookStep3: TRUE")
	return true
}
func (at *abstractTemplate) hookStep4() {
	fmt.Println("Default implementation hookStep4\n")
}

/*
CONCRETE BLUE CLASS (STRUCT) WITH SUB-METHODS IMPLEMENTATIONS
*/
type implementationBlue struct {
	abstractTemplate
}
func (ib *implementationBlue) abstractStep1() {
	fmt.Println("BLUE implementation abstractStep1")
}
func (ib *implementationBlue) abstractStep2() {
	fmt.Println("BLUE implementation abstractStep2")
}
//Optionally Overriding a Hook method
func (ib *implementationBlue) hookStep4() {
	fmt.Println("BLUE implementation hookStep4\n")
}
/*
CONCRETE RED CLASS (STRUCT) RED WITH SUB-METHODS IMPLEMENTATIONS
*/
type implementationRed struct {
	abstractTemplate
}
func (ir *implementationRed) abstractStep1() {
	fmt.Println("RED implementation abstractStep1")
}
func (ir *implementationRed) abstractStep2() {
	fmt.Println("RED implementation abstractStep2")
}
//Optionally Overriding a Hook method
func (ir *implementationRed) hookStep3() bool {
	fmt.Println("RED implementation hookStep3: FALSE (not executing step4)")
	return false
}


/*
CLIENT
 */
func main() {
	implementationBlue := &implementationBlue{}
	templateBlue := abstractTemplate{implementationBlue}
	templateBlue.templateMethod()

	implementationRed := &implementationRed{}
	templateRed := abstractTemplate{implementationRed}
	templateRed.templateMethod()
}