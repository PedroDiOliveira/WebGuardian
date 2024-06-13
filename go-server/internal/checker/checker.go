package checker

import "net/http"

type HTTPMessage struct {
	Code        int    `json:"code"`
	Description string `json:"desc"`
}

type errorMessage struct {
	Err string `json:"err"`
}

var messages = map[int]HTTPMessage{
	404: {Code: 404, Description: "Not found"},
	200: {Code: 200, Description: "Online Server"},
	429: {Code: 429, Description: "Too many requests"},
	408: {Code: 408, Description: "Request timeout"},
	500: {Code: 500, Description: "Internal Server Error"},
	502: {Code: 502, Description: "Bad gateway"},
	503: {Code: 503, Description: "Service unavailable"},
}

func Check(url string) (HTTPMessage, errorMessage) {
	resp, _ := http.Get(url)

	message := messages[resp.StatusCode]

	if message != nil {
		return message, nil
	}

	return _, 
}
