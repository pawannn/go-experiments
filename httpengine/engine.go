package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpEngine struct {
	mux *http.ServeMux
}

func NewHttpEngine() *HttpEngine {
	mux := http.NewServeMux()

	return &HttpEngine{
		mux: mux,
	}
}

func (e *HttpEngine) AddRoutes(routes []Route) {
	for _, route := range routes {
		var handler http.Handler = http.HandlerFunc(route.Controller)

		// middlewares are executed in reverse order
		for i := len(route.Middleware) - 1; i >= 0; i-- {
			handler = route.Middleware[i](handler)
		}

		finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != route.Method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte("Method Not Allowed"))
				return
			}

			ctx := SetContext(r.Context())
			r = r.WithContext(ctx)

			handler.ServeHTTP(w, r)
		})

		e.mux.Handle(route.Path, finalHandler)
	}
}

func (e *HttpEngine) SendResponse(w http.ResponseWriter, reqID string, statusCode int, clientMessage string, data interface{}) {
	w.WriteHeader(statusCode)
	response := Response{
		ReqID:         reqID,
		StatusCode:    statusCode,
		ClientMessage: clientMessage,
		Data:          data,
	}

	responseInBytes, _ := json.Marshal(response)
	w.Write(responseInBytes)
}

func (e *HttpEngine) Start(port uint16) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), e.mux)
	if err != nil {
		return err
	}

	return nil
}
