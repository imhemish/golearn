package main

import (
	"fmt"
	"sync"
	"time"
)

/* import (
	"fmt"
)

var chan1 = make(chan string)

func main() {
	go displayData()

	go addData("hemish")
	go addData("6")
	select {}

}

func displayData() {
	for {
		temp := <-chan1
		fmt.Println(temp)
	}
}

func addData(data string) {
	fmt.Println("go routine addData added", data)
	chan1 <- data
} */

var imageChan = make(chan string)
var wg sync.WaitGroup

func main() {
	// Random url
	images := []string{"sjdflskd", "ldsfj", "dkjf", "dadf", "asdfj", "adsf", "adfsdflkj", "aksdfjlsdjflsdj", "aksdfjslk", "werqer", "aqwero", "aerqpa", "aesrowj", "ajdflj", "qpulkr", "saldkjqp"}

	for {
		go func() {
			imageURL := <-imageChan
			time.Sleep(time.Second * 2)
			fmt.Printf("Image %v stored in database", imageURL)
		}()
	}

	select {}

}
