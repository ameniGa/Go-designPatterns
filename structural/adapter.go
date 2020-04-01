package main

import "fmt"

type adaptee struct {}

func (a *adaptee) specificRequest() {
	fmt.Println("the adaptee executes specificRequest")
}

// Itarget is used by the client to call the original function
// is the contract between the client and the adaptee
type Itarget interface {
	request()
}

type client struct {
	adapter Itarget
}

type adapter struct {
	adaptee adaptee
}

func (a adapter) request() {
	a.adaptee.specificRequest()
}

func main() {
	client := client{adapter:adapter{}}
	client.adapter.request()
}

