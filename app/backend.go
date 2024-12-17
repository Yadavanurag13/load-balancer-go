package main

import (
	"fmt"
	"log"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request)  {
	clientIp := r.RemoteAddr
	log.Printf("Recived request from %s\n%s %s %s\n", clientIp, r.Method, r.URL.Path, r.Proto)

	w.WriteHeader(http.StatusOK)
	message := "Hello From Backend server\n"
	fmt.Fprintf(w, message)

	log.Println("Replied with a hello message")
}
func main()  {
	http.HandleFunc("/", httpHandler)

	log.Printf("Starting Backend server on port 8081....\n")
	
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}