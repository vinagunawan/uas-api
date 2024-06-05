package main

import (
	"fmt"
	"uas-api/core"
	"uas-api/service/router"
)

func main() {
	app := core.NewApp()

	env := app.Env
	gin := app.Web

	router := router.RouterConstructor(gin, app)
	router.NewRouter()

	gin.Run(fmt.Sprintf(":%s", env.Port))
}