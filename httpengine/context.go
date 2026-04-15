package main

import (
	"context"

	"github.com/google/uuid"
)

type httpContext string

var RequestKey httpContext = "reqKey"

type Request struct {
	ReqID string `json:"req_id"`
}

func (e *HttpEngine) SetContext(ctx context.Context) context.Context {
	reqID := uuid.New().String()
	request := Request{
		ReqID: reqID,
	}

	c := context.WithValue(ctx, RequestKey, request)
	return c
}

func (e *HttpEngine) ParseContext(ctx context.Context) Request {
	val := ctx.Value(RequestKey)
	if val == nil {
		return Request{}
	}

	meta, ok := val.(Request)
	if !ok {
		return Request{}
	}

	return meta
}
