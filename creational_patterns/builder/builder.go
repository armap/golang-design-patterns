package main

import "fmt"

/*
PRODUCT
 */
type Product interface {
	doSomething()
}
type productOne struct {
	color string // required, set by builder type
	fieldA string // required, set by parameter
	fieldB string // optional, set by parameter
	fieldC string // optional, set by parameter
}
func (p *productOne) doSomething() {
	fmt.Println("builder doing something")
}

/*
PRODUCT BUILDERS
*/
type ProductBuilder interface {
	setA(string) ProductBuilder
	setOther(string) ProductBuilder
	getProduct() Product
	//reset()
}
func NewYellowProductOneBuilder() *yellowProductOneBuilder {
	return &yellowProductOneBuilder{}
}
type yellowProductOneBuilder struct {
	fieldA string
	fieldB string
}
func (ypb *yellowProductOneBuilder) setA(a string) ProductBuilder {
	ypb.fieldA = a
	return ypb
}
func (ypb *yellowProductOneBuilder) setOther(b string) ProductBuilder {
	ypb.fieldB = b
	return ypb
}
func (ypb *yellowProductOneBuilder) getProduct() Product {
	return &productOne{
		color: "yellow",
		fieldA: ypb.fieldA,
		fieldB: ypb.fieldB,
	}
}
//type greenProductOneBuilder struct {
//	fieldA string
//	fieldC string
//}

/*
DIRECTOR (optional)
*/
func NewDirector(b ProductBuilder) *director {
	return &director{ builder: b }
}
type director struct {
	builder ProductBuilder
}
func (d *director) BuildYellowProductOne() {
	d.builder.setA("A").setOther("B")
}
//func (d *director) BuildGreenProductOne() {
//	d.builder.setA("A").setOther("C")
//}


/*
CLIENT
*/
func main() {
	yellowProductOneBuilder := NewYellowProductOneBuilder()
	director := NewDirector(yellowProductOneBuilder)
	director.BuildYellowProductOne()
	yellowProduct := yellowProductOneBuilder.getProduct()
	yellowProduct.doSomething()
}



