package controllers

import (
	"github.com/labstack/echo/v4"
	"rizadwi.com/config"
	"rizadwi.com/models"
)

func PostArticles(c echo.Context) error {
	article := new(models.Article)
	if err := c.Bind(article); err != nil {
		return err
	}

	config.DB.Create(&article)

	return c.JSON(200, map[string]interface{}{
		"article": article,
		"message": "success",
	})
}

func DeleteArticles(c echo.Context) error {
	id := c.Param("id")
	article := models.Article{}
	result := config.DB.First(&article, id)

	if result.Error != nil {
		return c.JSON(404, map[string]interface{}{
			"message": "Article not found",
		})
	}

	config.DB.Delete(&article)

	return c.JSON(200, map[string]interface{}{
		"message":       "success",
		"deleteArticle": article,
	})
}
