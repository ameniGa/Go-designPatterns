package main

import (
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

type sumFunc func(int) int

func sum(n int) int {
	ch := make(chan int, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(ch chan int) {
			ch <- i
			wg.Done()
		}(ch)
	}
	defer close(ch)
	res := 0
	for nb := range ch {
		log.Println(nb)
		res += nb
	}
	return res
}

func loggerDecorator(sumFunc sumFunc, logger *log.Logger) sumFunc{
	return func(n int) int {
		fn := func(n int) (res int){
			defer func(t time.Time) {
				logger.Printf("it takes me %v to count the sum of %v values and the result is %v",time.Since(t), n , res)
			}(time.Now())
			return sumFunc(n)
		}
		return fn(n)
	}
}


func main() {
	//a := sum(100)
	//log.Println(a)
	f := loggerDecorator(sum,log.New(os.Stdout,"test",1))
	f(100)
	wg.Wait()
}
