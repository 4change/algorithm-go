package main

import (
	"fmt"
	"time"
)

var fooChan = make(chan struct{})
var barChan = make(chan struct{})

var n = 10

type FooBar struct {
	
}

func (f *FooBar) foo() {
	for i := 0; i < n; i++ {
		<- fooChan
		fmt.Printf("foo")
		barChan <- struct{}{}
	}
}

func (fb *FooBar) bar() {
	for i := 0; i < n; i++ {
		<- barChan
		fmt.Println("bar")
		fooChan <- struct{}{}
	}
}

func main() {
	fb := &FooBar{}

	go fb.bar()
	go fb.foo()
	
	fooChan <- struct{}{}

	time.Sleep(time.Second)
}