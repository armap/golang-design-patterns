package main

import "fmt"

/*
PRODUCT INTERFACE
 */
type Product interface {
	notDecoratedMethod() int
	decoratedMethod() string
}

/*
NORMAL PRODUCTS
*/
type productOne struct {}
func (po *productOne) notDecoratedMethod() int {
	return 1
}
func (po *productOne) decoratedMethod() string {
	return "product one"
}

type productTwo struct {}
func (pt *productTwo) notDecoratedMethod() int {
	return 2
}
func (pt *productTwo) decoratedMethod() string {
	return "product two"
}

/*
DECORATOR PRODUCTS
 */
type productDecorator struct {
	decoratedProduct Product
}
func (pd *productDecorator) notDecoratedMethod() int {
	return pd.decoratedProduct.notDecoratedMethod()
}

type productDecoratorBlue struct {
	productDecorator
}
func (pdb *productDecoratorBlue) decoratedMethod() string {
	return pdb.decoratedProduct.decoratedMethod() + " is Blue"
}

type productDecoratorSmall struct {
	productDecorator
}
func (pds *productDecoratorSmall) decoratedMethod() string {
	return pds.decoratedProduct.decoratedMethod() + " is Small"
}


/*
CLIENT
*/
func main() {
	productOne := &productOne{}
	fmt.Println(productOne.notDecoratedMethod())
	fmt.Println(productOne.decoratedMethod())

	//decorating blue over productOne
	productOneBlue := &productDecoratorBlue{
		productDecorator{decoratedProduct: productOne},
	}
	fmt.Println(productOneBlue.notDecoratedMethod())
	fmt.Println(productOneBlue.decoratedMethod())

	//decorating small over productOneBlue
	productOneBlueAndSmall := &productDecoratorSmall{
		productDecorator{decoratedProduct: productOneBlue},
	}
	fmt.Println(productOneBlueAndSmall.notDecoratedMethod())
	fmt.Println(productOneBlueAndSmall.decoratedMethod())
}