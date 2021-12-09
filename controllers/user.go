package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"rizadwi.com/config"
	"rizadwi.com/models"
)

func PostUsers(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	config.DB.Create(&user)

	return c.JSON(200, map[string]interface{}{
		"user":    user,
		"message": "success",
	})
}

func GetUsers(c echo.Context) error {
	var users []models.User
	// var user User

	config.DB.Find(&users)
	// db.Find(&user)

	fmt.Println("sending response...")
	return c.JSON(200, users)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	result := config.DB.First(&user, id)

	if result.Error != nil {
		return c.JSON(404, map[string]interface{}{
			"message": "User not found",
		})
	}

	config.DB.Delete(&user)

	return c.JSON(200, map[string]interface{}{
		"message":    "success",
		"deleteUser": user,
	})
}
