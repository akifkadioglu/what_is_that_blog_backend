package topiccontroller

import (
	"setup/database"
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
)

func GetAllTopic(c echo.Context) error {
	var topics []models.Topic
	db := database.DBManager()
	db.Joins("Author").Scopes(helpers.Paginate(c)).Order("created_at desc").Find(&topics)
	return c.JSON(200, map[string][]models.Topic{
		"topics": topics,
	})
}
