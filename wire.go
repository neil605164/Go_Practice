//go:build wireinject
// +build wireinject

package main

import (
	"Go_Practice/app/handler"

	"github.com/google/wire"
)

func InitializeEvent() handler.Event {
	wire.Build(handler.NewEvent, handler.NewGreeter, handler.NewMessage)
	return handler.Event{}
}
