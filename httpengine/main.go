package main

import (
	"fmt"
	"net/http"
)

type CoffeeHandler struct {
	e *HttpEngine
}

func (c *CoffeeHandler) GetCoffee(w http.ResponseWriter, r *http.Request) {
	meta := c.e.ParseContext(r.Context())

	clientMessage := "coffee prepared"
	response := map[string]string{
		"type": "star bucks",
	}

	c.e.SendResponse(w, meta.ReqID, http.StatusOK, clientMessage, response)
}

func (c *CoffeeHandler) GetCoffeeByID(w http.ResponseWriter, r *http.Request) {
	meta := c.e.ParseContext(r.Context())

	id := c.e.GetParams(r, "id")

	clientMessage := "coffee prepared"
	response := map[string]string{
		"type": "star bucks",
		"id":   id,
	}

	c.e.SendResponse(w, meta.ReqID, http.StatusOK, clientMessage, response)
}

func (c *CoffeeHandler) CoffeeMiddleware1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(1)
		h.ServeHTTP(w, r)
	})
}

func (c *CoffeeHandler) CoffeeMiddleware2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(2)
		h.ServeHTTP(w, r)
	})
}

func (c *CoffeeHandler) AddRoutes() {
	c.e.AddRoutes([]Route{
		{
			Method:      "GET",
			Path:        "/coffee",
			Description: "This route returns coffee response",
			Controller:  c.GetCoffee,
			Middleware: []Middleware{
				c.CoffeeMiddleware1,
				c.CoffeeMiddleware2,
			},
		},
		{
			Method:      "GET",
			Path:        "/coffee/{id}/",
			Description: "This route returns coffee id response",
			Controller:  c.GetCoffee,
			Middleware: []Middleware{
				c.CoffeeMiddleware1,
				c.CoffeeMiddleware2,
			},
		},
	})
}

const port uint16 = 8080

func main() {
	e := NewHttpEngine()

	coffeeHandler := CoffeeHandler{e}
	coffeeHandler.AddRoutes()

	fmt.Println("Server running at port :", port)
	e.Start(port)
}
