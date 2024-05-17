package handler

import (
	"log"
	"net/http"
	"time"
	"web-guardian/pkg/utils"
)

type Router struct{}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	utils.AllowCrude(res)

	path := req.URL.Path

	if path == "/check" && req.Method == "POST" {

	}
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
