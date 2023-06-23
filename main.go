package main

import (
	"fmt"
	"go-api/config"
	"go-api/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnetDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server running on port " + config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
