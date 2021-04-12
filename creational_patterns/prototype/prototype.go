package main

import (
	"fmt"
)

/*
PROTOTYPE INTERFACE
 */
type CloneableProduct interface {
	clone() CloneableProduct
	print()
}

/*
CONCRETE PROTOTYPES
*/
type product struct {
	fieldA string
}
type productOne struct {
	product
	fieldB string
	list   []int
}
func (po *productOne) clone() CloneableProduct {
	deepClone := *po

	deepClone.list = make([]int, len(po.list))
	copy(deepClone.list, po.list)

	return &deepClone
}
func (po *productOne) print() {
	fmt.Println("I'm a clone")
}

/*
REGISTRY
*/
var productPrototypes map[string]CloneableProduct

func LoadPrototypes() {
	productPrototypes = map[string]CloneableProduct{
		"one": &productOne{
			product{"A"},
			"B",
			[]int{1, 2, 3},
		},
	}
}
func ClonePrototype(productType string) CloneableProduct {
	return productPrototypes[productType].clone()
}


/*
CLIENT
*/
func main() {
	LoadPrototypes()
	clonedProduct := ClonePrototype("one")
	clonedProduct.print()
}



