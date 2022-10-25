package routes

import (
	"setup/controllers"
	contactcontroller "setup/controllers/ContactController"
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
	//Topics
	Network.GET("/searchTopics", topiccontroller.SearchTopics)
	Network.GET("/getAllTopics", topiccontroller.GetAllTopic)
	Network.GET("/getSingleTopic", topiccontroller.GetSingleTopic)

	//Contacts
	Network.POST("/createContact", contactcontroller.CreateContact)

}
