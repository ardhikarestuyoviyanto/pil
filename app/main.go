package main

import (
	"pil/config"
	"pil/database"
	"pil/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.InitConfiguration()
	e := echo.New()
	routes.AllRoutesAPI(e, database.InitDB(config), config)
	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
