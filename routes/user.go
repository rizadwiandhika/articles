package routes

import (
	"fmt"
	"math/rand"

	echo "github.com/labstack/echo/v4"
	"rizadwi.com/controllers"
)

func LoadUserRoutes(e *echo.Echo) {
	users := e.Group("/users")

	users.GET(
		"",
		controllers.GetUsers,
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				// fmt.Println("This is a middleware")
				// return hf(c)
				return c.JSON(200, map[string]interface{}{"sike": "ahahahah"})
			}
		})

	e.Use(func(hf echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("hello in use middleware 2")
		return func(c echo.Context) error {
			fmt.Println("hello in handler 2")

			c.Set("request", c.Request())
			return hf(c)
		}
	})

	users.GET("/test",
		func(c echo.Context) error {

			return c.JSON(200, map[string]interface{}{
				"message": "Hello World",
				"pesan":   c.Get("pesan"),
				"token":   c.Get("token"),
			})
		},
		func(hf echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Set("pesan", "apa kabar")

				if rand.Float64() > 0.5 {
					c.Set("break", true)
				}

				return hf(c)
			}
		},
		func(hf echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				if c.Get("break") == true {
					return hf(c)
				}

				c.Set("token", "qejoqim,qnwekq")
				return hf(c)
			}
		},
	)

	users.POST("", controllers.PostUsers)
	users.DELETE("/:id", controllers.DeleteUser)
}

/*
e.GET("/users", GetUsers)
e.POST("/users", PostUsers)
e.DELETE("/users/:id", DeleteUser)
*/
