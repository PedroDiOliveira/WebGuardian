package model

type URLrequest struct {
	URL string `json:"url"`
}

type URLresponse struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}
