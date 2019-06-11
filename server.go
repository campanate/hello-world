package main

import (
	"io"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		HealthCheckHandler(w, r)
	})

	listener, err := net.Listen("tcp", "0.0.0.0:8080")

	if(err != nil){
		fmt.Errorf("Error: %s", err.Error())
	}

	log.Print("Serving server in localhost 8080...")
	err = http.Serve(listener, nil)
}


// HealthCheckHandler check if the server is health
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
   
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")


    io.WriteString(w, `{"health": true}`)
}
