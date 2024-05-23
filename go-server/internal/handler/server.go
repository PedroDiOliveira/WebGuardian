package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"web-guardian/internal/checker"
	"web-guardian/pkg/utils"
)

type Router struct{}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// Adicionando cabeçalhos CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	path := req.URL.Path

	if path == "/check" && req.Method == "POST" {
		url := utils.GetUrl(req.Body)
		response := checker.Check(url)

		body, _ := json.Marshal(response)
		res.Write(body)
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
