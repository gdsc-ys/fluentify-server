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
	authMiddlewareImpl := middleware.AuthMiddlewareInit(client)
	userServiceImpl := service.UserServiceInit(client)
	storageClient := NewFirebaseStorageClient(app)
	storageServiceImpl := service.StorageServiceInit(storageClient)
	firestoreClient := NewFireStoreClient(app)
	topicServiceImpl := service.TopicServiceInit(firestoreClient)
	sentenceServiceImpl := service.SentenceServiceInit(firestoreClient)
	sceneServiceImpl := service.SceneServiceInit(firestoreClient)
	userHandlerImpl := handler.UserHandlerInit(userServiceImpl)
	topicHandlerImpl := handler.TopicHandlerInit(topicServiceImpl, storageServiceImpl)
	sentenceHandlerImpl := handler.SentenceHandlerInit(sentenceServiceImpl, storageServiceImpl)
	sceneHandlerImpl := handler.SceneHandlerInit(sceneServiceImpl, storageServiceImpl)
	feedbackHandlerImpl := handler.FeedbackHandlerInit(sentenceServiceImpl, sceneServiceImpl, storageServiceImpl)
	initialization := NewInitialization(authMiddlewareImpl, userServiceImpl, storageServiceImpl, topicServiceImpl, sentenceServiceImpl, sceneServiceImpl, userHandlerImpl, topicHandlerImpl, sentenceHandlerImpl, sceneHandlerImpl, feedbackHandlerImpl)
	return initialization
}

// wire.go:

var firebaseApp = wire.NewSet(InitializeFirebaseApp)

var firebaseAuthClient = wire.NewSet(NewFirebaseAuthClient)

var firebaseStorageClient = wire.NewSet(NewFirebaseStorageClient)

var fireStoreClient = wire.NewSet(NewFireStoreClient)

var authMiddlewareSet = wire.NewSet(middleware.AuthMiddlewareInit, wire.Bind(new(middleware.AuthMiddleware), new(*middleware.AuthMiddlewareImpl)))

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var storageServiceSet = wire.NewSet(service.StorageServiceInit, wire.Bind(new(service.StorageService), new(*service.StorageServiceImpl)))

var topicServiceSet = wire.NewSet(service.TopicServiceInit, wire.Bind(new(service.TopicService), new(*service.TopicServiceImpl)))

var sentenceServiceSet = wire.NewSet(service.SentenceServiceInit, wire.Bind(new(service.SentenceService), new(*service.SentenceServiceImpl)))

var sceneServiceSet = wire.NewSet(service.SceneServiceInit, wire.Bind(new(service.SceneService), new(*service.SceneServiceImpl)))

var userHandlerSet = wire.NewSet(handler.UserHandlerInit, wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)))

var topicHandlerSet = wire.NewSet(handler.TopicHandlerInit, wire.Bind(new(handler.TopicHandler), new(*handler.TopicHandlerImpl)))

var sentenceHandlerSet = wire.NewSet(handler.SentenceHandlerInit, wire.Bind(new(handler.SentenceHandler), new(*handler.SentenceHandlerImpl)))

var sceneHandlerSet = wire.NewSet(handler.SceneHandlerInit, wire.Bind(new(handler.SceneHandler), new(*handler.SceneHandlerImpl)))

var feedbackHandlerSet = wire.NewSet(handler.FeedbackHandlerInit, wire.Bind(new(handler.FeedbackHandler), new(*handler.FeedbackHandlerImpl)))
