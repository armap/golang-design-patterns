package main

import (
	"fmt"
	"sync"
)


type singleton struct { /* fields */  }

var singletonInstance *singleton

var once sync.Once

func getInstance() *singleton {
	once.Do(func ()  {
		if singletonInstance == nil {
			singletonInstance = &singleton{ /* data */ }
		}
	})
	return singletonInstance
}

func (p *singleton) doSomething() {
	fmt.Println("singleton doing something")
}

/*
CLIENT
*/
func main() {
	singletonInstance := getInstance()
	singletonInstance.doSomething()
}
