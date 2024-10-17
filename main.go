package main

import (
	"exercise-app/routes"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
  e.Use(middleware.Static("public"))

	routes.INDEX.Build(e)
  routes.AUTH.Build(e)

	log.Fatal(e.Start(os.Getenv("GO_PORT")))
}
