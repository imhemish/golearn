package main

import (
	"fmt"
	"time"
)

const MAX_LIMIT = 2

func main() {
	limiter := make(chan struct{}, MAX_LIMIT)
	imagesToProcess := []string{}
	for i := 1; i < 20; i++ {
		imagesToProcess = append(imagesToProcess, fmt.Sprint(i))
	}
	for _, image := range imagesToProcess {
		go processImage(image, &limiter)
	}

	select {}

}

func processImage(image string, limiter *chan struct{}) {
	*limiter <- struct{}{}
	println("acquired limiter")
	println("processing image")
	time.Sleep(time.Second * 2)
	<-*limiter
	println("limiter released")
}
