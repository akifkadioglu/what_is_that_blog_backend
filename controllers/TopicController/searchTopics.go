package topiccontroller

import (
	"setup/database"
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
)

func SearchTopics(c echo.Context) error {
	var topics []models.Topic
	input := c.QueryParam("title")
	db := database.DBManager()

	db.Joins("Author").Where(`title LIKE "%` + input + `%"`).Scopes(helpers.Paginate(c)).Order("title").Find(&topics)
	return c.JSON(200, map[string][]models.Topic{
		"topics": topics,
	})
}
