package main 

import (
    "fmt"
    "sync"
)

func publisher(wg *sync.WaitGroup, msgChan chan string) {

    for i := range 10 { 
	msgChan <- fmt.Sprintf("Message: %d", i) 
    } 
    close(msgChan) 
    wg.Done()
}


func subscriber(id int, wg *sync.WaitGroup, msgChan chan string) {

    
    for message := range msgChan {
	fmt.Printf("Subscriber %d received message: %s\n", id, message) 
    } 

    wg.Done()

}


func main() {

     
    var msgChan = make(chan string) 

    wg := &sync.WaitGroup{} 

    wg.Add(1)

    go publisher(wg, msgChan) 

    for i := range 3 {
	wg.Add(1) 
	go subscriber(i, wg, msgChan) 

    } 

    wg.Wait() 
}
