package main

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Middleware func(http.Handler) http.Handler

type Route struct {
	Method      string
	Path        string
	Description string
	Controller  HandlerFunc
	Middleware  []Middleware
}

type Response struct {
	ReqID         string      `json:"req_id"`
	StatusCode    int         `json:"status_cdoe"`
	ClientMessage string      `json:"client_message"`
	Data          interface{} `json:"data"`
}
