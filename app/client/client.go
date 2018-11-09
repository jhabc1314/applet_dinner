package client

import "net/http"

var GetClient *http.Client

var PostClient *http.Client

func init() {
	GetClient = &http.Client{}
	PostClient = &http.Client{}
}
