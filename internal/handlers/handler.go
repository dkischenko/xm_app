package handlers

import (
	"github.com/gorilla/mux"
)

//go:generate mockgen -source=handler.go -destination=mocks/handler_mock.go
type Handler interface {
	Register(router *mux.Router)
}
