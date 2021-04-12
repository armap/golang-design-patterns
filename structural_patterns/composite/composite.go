package main

import (
	"fmt"
	"strconv"
)

/*
COMPONENT INTERFACE
 */
type component interface {
	Description() string
	Total() int
}

/*
LEAF
 */
type leaf struct {
	name string
	total int
}
func (l *leaf) Description() string {
	return l.name + " - " + strconv.Itoa(l.total) + "\n"
}
func (l *leaf) Total() int {
	return l.total
}

/*
COMPOSITE
*/
type composite struct {
	components []component
}
func (c *composite) Add(component component) {
	c.components = append(c.components, component)
}
func (c *composite) Description() string {
	var description string
	for _, component := range c.components {
		description += component.Description()

	}
	return description
}
func (c *composite) Total() int {
	total := 0
	for _, component := range c.components {
		total += component.Total()
	}
	return total
}


/*
CLIENT
*/
func main() {
	multipleOrderA := new(composite)
	multipleOrderA.Add(&leaf{name: "Order1", total: 10})
	multipleOrderA.Add(&leaf{name: "Order2", total: 20})
	multipleOrderA.Add(&leaf{name: "Order3", total: 30})

	multipleOrderB := new(composite)
	multipleOrderB.Add(multipleOrderA)
	multipleOrderB.Add(&leaf{name: "Order4", total: 40})

	fmt.Println(multipleOrderB.Description())
	fmt.Printf("Total: %d", multipleOrderB.Total())
}