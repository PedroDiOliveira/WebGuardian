package utils

import "net/http"

func AllowCrude(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
