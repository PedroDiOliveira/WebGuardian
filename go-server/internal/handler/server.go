package handler

import (
	"log"
	"net/http"
	"time"
)

type Router struct{}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

}

func RunServer() {
	s := http.Server{
		Addr:         ":8080",
		Handler:      Router{},
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
