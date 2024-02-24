package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var limiter = make(chan struct{}, 4)
var processWG sync.WaitGroup

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		println("server started")
		for i := 1; i <= 30; i++ {
			processWG.Add(1)
			go processImage(fmt.Sprint(rand.Int()), ctx, &processWG)
		}
	}()

	select {
	case <-ctx.Done():
		println("Quitting server")
		time.Sleep(time.Second * 2)
		// Server shutdown code here
		println("SErver shutdown done")
	}
}

func processImage(sourceURL string, parent_context context.Context, wg *sync.WaitGroup) {
	var task_completed = make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(parent_context)
	defer cancel()
	limiter <- struct{}{}
	fmt.Printf("processing image %v\n", sourceURL)
	time.Sleep(time.Second * 3)
	<-limiter
	task_completed <- struct{}{}
	wg.Done()

	select {
	case <-ctx.Done():
		println("Task Cancelled %v\n", sourceURL)
	case <-task_completed:
		fmt.Printf("task done %v\n", sourceURL)

	}

}
