package main

import (
	"fmt"
	"sync"
)

/*
SHARED (INTRINSIC) FLYWEIGHT OBJECT
*/
type sharedFlyweight struct {
	name string
	color string
}
func newSharedFlyweight(name, color string) *sharedFlyweight {
	return &sharedFlyweight{name, color}
}

/*
UNIQUE (EXTRINSIC) FLYWEIGHT OBJECT
*/
type uniqueFlyweight struct {
	sharedFlyweight *sharedFlyweight
	fieldA          string
	fieldB          string
}
func newUniqueFlyweight(sharedFlyweight *sharedFlyweight, fieldA, fieldB string) *uniqueFlyweight {
	return &uniqueFlyweight{
		sharedFlyweight: sharedFlyweight,
		fieldA: fieldA,
		fieldB: fieldB,
	}
}
func (p *uniqueFlyweight) printFields() {
	fmt.Printf("Name: %s - Color: %s - Field A: %s - Field B: %s\n",
		p.sharedFlyweight.name,
		p.sharedFlyweight.color,
		p.fieldA,
		p.fieldB)
}

/*
FLYWEIGHT FACTORY
 */
var (
	// Map of pointers to share memory address
	sharedFlyweightCache map[string]*sharedFlyweight
	uniqueFlyweights []*uniqueFlyweight
	once sync.Once
)
func init() {
	once.Do(func() {
		sharedFlyweightCache = make(map[string]*sharedFlyweight)
	})
}
// Factory method to get a sharedFlyweight from CACHE or create a new one
func sharedFlyweightFactory(name, color string) *sharedFlyweight {
	if sharedFlyweightCache[name] == nil {
		sharedFlyweightCache[name] = newSharedFlyweight(name, color)
	}
	return sharedFlyweightCache[name]
}
func AddFlyweight(name, color, fieldA, fieldB string) {
	sharedFlyweight := sharedFlyweightFactory(name, color)
	uniqueFlyweight := newUniqueFlyweight(sharedFlyweight, fieldA, fieldB)
	uniqueFlyweights = append(uniqueFlyweights, uniqueFlyweight)
	uniqueFlyweight.printFields()
	return
}


/*
CLIENT
 */
func main() {
	AddFlyweight("First", "blue", "big", "rounded")
	AddFlyweight("Second","red", "small", "squared")
}