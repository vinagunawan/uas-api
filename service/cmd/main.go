package main

import (
	"backend/core"
	"backend/service/router"

	"github.com/gorilla/mux"
)

func main() {
	core.LoadEnv()
	core.ConnectDatabase()

	r := mux.NewRouter()
	router.RegisterRoutes(r)

	core.RunServer(r)
}
