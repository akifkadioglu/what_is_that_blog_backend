package topiccontroller

import (
	"setup/database"
	"setup/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetSingleTopic(c echo.Context) error {
	var topic models.Topic
	db := database.DBManager()
	id, _ := strconv.Atoi(c.QueryParam("id"))
	db.Joins("Author").First(&topic, id)
	return c.JSON(200, map[string]models.Topic{
		"topic": topic,
	})
}
