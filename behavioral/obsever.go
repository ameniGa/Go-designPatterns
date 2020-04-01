package main

import (
	"log"
	"sync"
)

type event struct {
	data int
}

/*****************************************************************/

type observer interface {
	update(event event)
}

/*****************************************************************/

type observable interface {
	subscribe(observer observer)
	unsubscribe(observer observer)
	notify(event event)
}

/*****************************************************************/

type publisher struct {
	observers sync.Map
}

func (p *publisher) subscribe(observer observer) {
	p.observers.Store(observer, struct {}{})
	log.Printf("observer: %+v have just subscribed",observer)
}

func (p *publisher) unsubscribe(observer observer) {
	_, ok := p.observers.Load(observer)
	if ok {
		p.observers.Delete(observer)
		log.Printf("observer: %+v have just unsubscribed",observer)
	}
}

func (p *publisher) notify(event event) {
	p.observers.Range(func(key, value interface{}) bool {
		log.Printf("notify the observer: %+v",key)
		key.(observer).update(event)
		return true
	})
}

/*****************************************************************/

type subscriber struct {
	id int
}

func (o subscriber) update(event event) {
	log.Printf("observer: %v, received: %v",o.id, event.data)
}

/*****************************************************************/

func main() {
	subscriber := subscriber{id: 10}
	var publisher publisher
	publisher.subscribe(subscriber)
	publisher.notify(event{data:100})
	//time.Sleep(3*time.Second)
	publisher.unsubscribe(subscriber)
}
