package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type HttpEngine struct {
	router chi.Router
}

func NewHttpEngine() *HttpEngine {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return &HttpEngine{
		router: router,
	}
}

func (e *HttpEngine) GetParams(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func (e *HttpEngine) GetQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func (e *HttpEngine) Start(port uint16) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), e.router)
	if err != nil {
		return err
	}

	return nil
}
