package routes

import (
	"setup/controllers"
	topiccontroller "setup/controllers/TopicController"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Web() {

	Network.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	Network.GET("/index", controllers.Index)
	//Home Page
	Network.GET("/getAllTopics", topiccontroller.GetAllTopic)
	Network.GET("/getSingleTopic", topiccontroller.GetSingleTopic)
	//forgot password items
	//it doesnt work now
	/* Network.GET("/accept-password", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	})
	Network.POST("/change-password", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}) */
}
