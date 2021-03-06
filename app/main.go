package main

import (
	"app/resolver"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.POST("/graphql", func(c echo.Context) error {
		config := resolver.Config{
			Resolvers: resolver.New(),
		}
		h := handler.GraphQL(resolver.NewExecutableSchema(config))
		h.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":3000"))
}
