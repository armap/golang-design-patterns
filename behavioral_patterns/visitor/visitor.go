package main

import "fmt"
/**
VISITOR pattern in conjunction with COMPOSITE pattern
*/


/*
ELEMENT INTERFACE
*/
type Element interface {
	accept(Visitor)
	Quality() string
}
/*
CONCRETE ELEMENT ONE
 */
type elementOne struct {
	quality string
}
func NewElementOne() *elementOne {
	return &elementOne{"One quality is High"}
}
func (eo *elementOne) accept(visitor Visitor) {
	visitor.visitOne(eo)
}
func (eo *elementOne) Quality() string {
	return eo.quality
}
/*
CONCRETE ELEMENT TWO
*/
type elementTwo struct {
	quality string
}
func NewElementTwo() *elementTwo {
	return &elementTwo{"Two quality is Low"}
}
func (et *elementTwo) accept(visitor Visitor) {
	visitor.visitTwo(et)
}
func (et *elementTwo) Quality() string {
	return et.quality
}
/*
CONCRETE ELEMENT ORDER
*/
type elementOrder struct {
	elements []Element
	averageQuality string
}
func (eor *elementOrder) accept(visitor Visitor) {
	for _, element := range eor.elements {
		element.accept(visitor)
	}
	visitor.visitOrder(eor)
}
func (eor *elementOrder) Quality() string {
	fmt.Println("Calculating average quality of all ElementOrders")
	averageQuality := "Average quality is Medium"
	return averageQuality
}
func (eor *elementOrder) addElement(e Element) {
	eor.elements = append(eor.elements, e)
}
func (eor *elementOrder) getElements() []Element {
	return eor.elements
}

/*
ELEMENT VISITOR INTERFACE
 */
type Visitor interface {
	visitOne(*elementOne)
	visitTwo(*elementTwo)
	visitOrder(*elementOrder)
}
/*
CONCRETE PRICE VISITOR
*/
type priceVisitor struct {
	price int
}
func (pv *priceVisitor) visitOne(eo *elementOne) {
	// Calculate price for elementOne.
	// Then assign in to the price instance variable.
	fmt.Println("Calculating price for element One")
	pv.price += 1
}
func (pv *priceVisitor) visitTwo(et *elementTwo) {
	fmt.Println("Calculating price for elementTwo")
	pv.price += 2
}
func (pv *priceVisitor) visitOrder(elementOrder *elementOrder) {
	if pv.price > 5 { pv.price -= 1 } //discount
	fmt.Printf("Total price is  %d\n\n", pv.price)
}
/*
CONCRETE DISPLAY VISITOR
*/
type displayVisitor struct {
	displaying string
}

func NewDisplayVisitor() *displayVisitor {
	return &displayVisitor{"Displaying: "}
}
func (dv *displayVisitor) visitOne(eo *elementOne) {
	fmt.Printf("%s%s\n", dv.displaying, eo.Quality())
}
func (dv *displayVisitor) visitTwo(et *elementTwo) {
	fmt.Printf("%s%s\n", dv.displaying, et.Quality())
}
func (dv *displayVisitor) visitOrder(elementOrder *elementOrder) {
	fmt.Printf("%s%s\n", dv.displaying, elementOrder.Quality())
}


/*
CLIENT
 */
func main() {
	//CREATING ELEMENTS
	elementOne := NewElementOne()
	elementTwo := NewElementTwo()
	elementOrder := new(elementOrder)
	elementOrder.addElement(elementOne)
	elementOrder.addElement(elementTwo)

	//PRICE VISITOR
	priceVisitor := new(priceVisitor)
	//ELEMENTS ACCEPTING PRICE VISITOR
	elementOrder.accept(priceVisitor)

	//DISPLAY VISITOR
	displayVisitor := NewDisplayVisitor()
	//ELEMENTS ACCEPTING DISPLAY VISITOR
	elementOrder.accept(displayVisitor)
}