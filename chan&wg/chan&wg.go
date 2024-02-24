package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var chan1 = make(chan string)
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go displayData(&wg)
	wg.Add(1)
	go feedData(&wg)
	wg.Wait()
}

func displayData(wg *sync.WaitGroup) {
	for i := range chan1 {
		fmt.Println(i)
	}
	wg.Done()
}

func feedData(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		chan1 <- fmt.Sprint(rand.Int())
	}
	close(chan1)
	wg.Done()

}
