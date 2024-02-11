// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"github.com/gdsc-ys/fluentify-server/src/handler"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"github.com/gdsc-ys/fluentify-server/src/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Init() *Initialization {
	app := InitializeFirebaseApp()
	client := NewFirebaseAuthClient(app)
	authMiddleware := middleware.AuthMiddleware{
		authClient: client,
	}
	userServiceImpl := service.UserServiceInit(client)
	userHandlerImpl := handler.UserHandlerInit(userServiceImpl)
	initialization := NewInitialization(authMiddleware, userServiceImpl, userHandlerImpl)
	return initialization
}

// wire.go:

var firebaseApp = wire.NewSet(InitializeFirebaseApp)

var firebaseAuthClient = wire.NewSet(NewFirebaseAuthClient)

var authMiddlewareSet = wire.NewSet(wire.Struct(new(middleware.AuthMiddleware), "*"))

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var userHandlerSet = wire.NewSet(handler.UserHandlerInit, wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)))
