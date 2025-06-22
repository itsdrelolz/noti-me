package main 

import (
    "fmt"
    "net/http"
    "log"

    "github.com/gorilla/websocket"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    
    w.Write([]byte("Gorilla!\n"))

}


func main() {
    r := mux.New

}
