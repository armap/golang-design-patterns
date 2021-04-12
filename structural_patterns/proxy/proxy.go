package main

import "fmt"

/*
SERVER INTERFACE
 */
type server interface {
	handleRequest(int)
}

/*
REAL SERVER
*/
type realServer struct {}
func (a *realServer) handleRequest(num int) {
	fmt.Printf("Request allowed for number: %d\n", num)
}

/*
PROXY SERVER
*/
type proxyServer struct {
	realServer *realServer
	allowed bool
}
func (n *proxyServer) handleRequest(num int) {
	n.accessControl(num)
	if n.allowed {
		n.realServer.handleRequest(num)
	} else {
		fmt.Printf("Request Not allowed for number: %d\n", num)
	}
}
func (n *proxyServer) accessControl(num int)  {
	if num < 10 {
		n.allowed = true
	} else {
		n.allowed = false
	}
}
func NewProxyServer() *proxyServer {
	return &proxyServer{
		realServer: new(realServer),
		allowed: false,
	}
}


/*
CLIENT
 */
func main() {
	proxyServer := NewProxyServer()
	proxyServer.handleRequest(11)
	proxyServer.handleRequest(8)
}