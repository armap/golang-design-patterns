package main

import "fmt"

/*
CONTEXT
*/
type Context struct {
	strategy    Strategy
}
func (c *Context) setStrategy(s Strategy) {
	c.strategy = s
}
func (c *Context) executeStrategy() {
	c.strategy.execute(c)
}

/*
STRATEGY INTERFACE
 */
type Strategy interface {
	execute(c *Context)
}
/*
CONCRETE STRATEGIES
 */
type strategyOne struct {
}
func (l *strategyOne) execute(c *Context) {
	fmt.Println("Executing Strategy One")
}
type strategyTwo struct {
}
func (l *strategyTwo) execute(c *Context) {
	fmt.Println("Executing Strategy Two")
}
type strategyThree struct {
}
func (l *strategyThree) execute(c *Context) {
	fmt.Println("Executing Strategy Three")
}


/*
CLIENT
 */
func main() {
	context := new(Context)

	context.setStrategy(new(strategyOne))
	context.executeStrategy()

	context.setStrategy(new(strategyTwo))
	context.executeStrategy()

	context.setStrategy(new(strategyThree))
	context.executeStrategy()
}