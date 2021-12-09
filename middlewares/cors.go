package middlewares

import "github.com/labstack/echo/v4"

func CORSMiddleware(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Method", "GET, POST, DELETE, PUT")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		return hf(c)
	}
}
