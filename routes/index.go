package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"rizadwi.com/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// * Di Echo, use itu dieksekusi awal2, jadi walauppun definisiin
	// * Use di akhir routes, si use akan tetep dieksekusi paling awal
	e.Use(middlewares.CORSMiddleware)

	e.Use(func(hf echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("hello in use middleware 1")
		return func(c echo.Context) error {
			fmt.Println("hello in handler 1")

			c.Set("request", c.Request())
			return hf(c)
		}
	})

	LoadUserRoutes(e)
	LoadArticleRoutes(e)

	e.Use(func(hf echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("hello in use middleware 3")
		return func(c echo.Context) error {
			fmt.Println("hello in handler 3")

			c.Set("request", c.Request())
			return hf(c)
		}
	})

	return e
}

/* e.GET("/users", GetUsers)
e.POST("/users", PostUsers)
e.DELETE("/users/:id", DeleteUser) */

/* e.POST("/articles", PostArticles)
e.DELETE("/articles/:id", DeleteArticles) */
