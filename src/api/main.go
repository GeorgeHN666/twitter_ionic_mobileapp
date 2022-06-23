package main

import (
	"log"

	"github.com/GeorgeHN666/twitter_ionic_mobileapp/src/api/db"
	"github.com/GeorgeHN666/twitter_ionic_mobileapp/src/api/routes"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Lost connection with database")
	}

	routes.Serve()
	log.Println("Successful connectio with database")

}
