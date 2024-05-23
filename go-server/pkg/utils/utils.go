package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"web-guardian/internal/model"
)

func AllowCrude(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

}

func GetUrl(body io.ReadCloser) string {
	var url model.URLrequest
	json.NewDecoder(body).Decode(&url)
	return url.URL
}
