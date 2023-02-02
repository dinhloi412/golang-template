package main

import (
	"log"

	"project/api/v1"

	"project/controller"

	"github.com/asdine/storm/v3"
)

func main() {

	// Create controller]
	port := ":5000"
	db, err := storm.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctrl := controller.NewController(port, db)

	// Create api server
	apiServer, err := controller.NewApiServer(ctrl.Port())
	if err != nil {
		log.Fatalln(err)
	}
	api := api.NewApi(ctrl)

	// Run server
	api.Run(apiServer)
}
