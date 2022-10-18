package contactcontroller

import (
	"setup/database"
	"setup/models"

	"github.com/labstack/echo/v4"
)

func CreateContact(c echo.Context) error {
	var contact models.Contact
	db := database.DBManager()
	c.Bind(&contact)
	result := db.Create(&contact)
	if result.Error != nil {
		return c.JSON(500, map[string]bool{
			"isCreated": false,
		})
	}
	return c.JSON(200, map[string]bool{
		"isCreated": true,
	})
}
