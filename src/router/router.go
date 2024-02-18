package router

import (
	"github.com/gdsc-ys/fluentify-server/config"
	"github.com/gdsc-ys/fluentify-server/src/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Router(init *config.Initialization) *echo.Echo {
	e := echo.New()

	e.Debug = true

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/GetUser", init.UserHandler.GetUser)
	e.POST("/UpdateUser", init.UserHandler.UpdateUser)
	e.POST("/DeleteUser", init.UserHandler.DeleteUser)

	e.POST("/ListTopics", init.TopicHandler.ListTopics)
	e.POST("/GetTopic", init.TopicHandler.GetTopic)

	e.POST("/GetSentence", init.SentenceHandler.GetSentence)
	e.POST("/GetScene", init.SceneHandler.GetScene)

	return e
}
