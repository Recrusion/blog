package handlers

import "github.com/Recrusion/blog-api/internal/service"

type Handlers struct {
	handler *service.Service
}

func NewHandlers(handler *service.Service) *Handlers {
	return &Handlers{
		handler: handler,
	}
}
