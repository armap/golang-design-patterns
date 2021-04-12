package main

import "fmt"

/*
FACTORY METHOD
*/
func GetProduct(productType int) Product {
	switch productType {
	case 1:
		return newProduct1()
	case 2:
		return newProduct2()
	}
	return nil
}

/*
ABSTRACT PRODUCT
 */
type product struct{}
type Product interface {
	prodMethod()
}

/*
CONCRETE PRODUCT 1
 */
func newProduct1() Product { return &product1{} }
type product1 struct{ product }
func (p1 *product1) prodMethod() { fmt.Println("product1") }

/*
CONCRETE PRODUCT 2
*/func newProduct2() Product { return &product2{} }
type product2 struct{ product }
func (p1 *product2) prodMethod() { fmt.Println("product2") }


/*
CLIENT
*/
func main() {
	product1 := GetProduct(1)
	product1.prodMethod()

	product2 := GetProduct(2)
	product2.prodMethod()
}
