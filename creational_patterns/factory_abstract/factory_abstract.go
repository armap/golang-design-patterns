package main

import "fmt"

/*
ABSTRACT-FACTORY METHOD
*/
func GetFactory(factoryType string) Factory {
	switch factoryType {
	case "A":
		return newFactoryA()
	case "B":
		return newFactoryB()
	}
	return nil
}

/*
ABSTRACT FACTORY
*/
type factory struct{}
type Factory interface {
	makeProduct1() Product1
	makeProduct2() Product2
}

/*
ABSTRACT PRODUCTS
*/
type product1 struct{}
type Product1 interface {
	prod1Method()
}
type product2 struct{}
type Product2 interface {
	prod2Method()
}

/*
CONCRETE FACTORY A AND ITS CONCRETE PRODUCTS 1A AND 2A
*/
func newFactoryA() Factory { return &factoryA{} }
type factoryA struct{ factory }

func (fA *factoryA) makeProduct1() Product1 { return &product1A{} }
type product1A struct{ product1 }
func (p1A *product1A) prod1Method() { fmt.Println("product1A") }

func (fA *factoryA) makeProduct2() Product2 { return &product2A{} }
type product2A struct{ product2 }
func (p2A *product2A) prod2Method() { fmt.Println("product2A") }

/*
CONCRETE FACTORY B AND ITS CONCRETE PRODUCTS 1B AND 2B
*/
func newFactoryB() Factory { return &factoryA{} }
type factoryB struct{ factory }

func (fB *factoryB) makeProduct1() Product1 { return &product1B{} }
type product1B struct{ product1 }
func (p1B *product1B) prod1Method() { fmt.Println("product1B") }

func (fB *factoryB) makeProduct2() Product2 { return &product2B{} }
type product2B struct{ product2 }
func (p2B *product2B) prod2Method() { fmt.Println("product2B") }


/*
CLIENT
*/
func main() {
	factoryA := GetFactory("A")
	product1A := factoryA.makeProduct1()
	product1A.prod1Method()
	product2A := factoryA.makeProduct2()
	product2A.prod2Method()

	factoryB := GetFactory("B")
	product1B := factoryB.makeProduct1()
	product1B.prod1Method()
	product2B := factoryB.makeProduct2()
	product2B.prod2Method()
}

