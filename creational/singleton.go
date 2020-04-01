package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// should be unexposed
type singleton struct{}

// should be unexposed
var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	if instance == nil {
		log.Printf("instance not created yet")
		once.Do(
			func() {
				log.Printf("create instance once")
				instance = &singleton{}
			})
	}
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			s2 := GetInstance()
			fmt.Printf("instance n°%v: %p\n",i, s2)
		}(i)
	}
	time.Sleep(5 * time.Second)
}
