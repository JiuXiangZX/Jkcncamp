package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	VERSION = "VERSION"
)

func main () {
	http.HandleFunc("/svc", WebHandler)
	http.HandleFunc("/healthz", HealthHandler)
	log.Print("http server is running")
	err := http.ListenAndServe("192.168.124.34:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal("http server run over")
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	rHeader := r.Header
	for k, v := range rHeader {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	w.Header().Add(VERSION, os.Getenv(VERSION))
	w.WriteHeader(http.StatusOK)
	fmt.Printf("client host:%s, response code:%d", r.Host, http.StatusOK)
}

func HealthHandler(w http.ResponseWriter, r * http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}