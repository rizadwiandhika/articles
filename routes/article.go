package routes

import (
	echo "github.com/labstack/echo/v4"
	"rizadwi.com/controllers"
)

func LoadArticleRoutes(e *echo.Echo) {
	articles := e.Group("/articles")

	articles.POST("", controllers.PostArticles)
	articles.DELETE("/:id", controllers.DeleteArticles)
}

/* e.POST("/articles", PostArticles)
e.DELETE("/articles/:id", DeleteArticles) */
