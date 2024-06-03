package core

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer(router *mux.Router) {
	port := GetEnv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
