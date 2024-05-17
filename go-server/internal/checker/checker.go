package checker

import "net/http"

type HTTPMessage struct {
	Code        int    `json:"code"`
	Description string `json:"desc"`
}

var messages = map[int]HTTPMessage{
	404: {Code: 404, Description: "Not found"},
	200: {Code: 200, Description: "OK"},
}

func Check(url string) HTTPMessage {
	resp, _ := http.Get(url)

	message := messages[resp.StatusCode]

	return message
}
