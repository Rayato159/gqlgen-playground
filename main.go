package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Rayato159/gqlgen-practice/graph"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	queryHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	e.POST("/query", func(c echo.Context) error {
		queryHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
