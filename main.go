package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahmedkhalaf1996/GRQGoTaskx/conf"
	"github.com/ahmedkhalaf1996/GRQGoTaskx/database"
	"github.com/ahmedkhalaf1996/GRQGoTaskx/router"
)

const defaultPort = "8080"

func main() {
	// -------- config ------------ //
	conf.CallConf()
	// ---------- confing ------------ //
	// ----- db --------------- //
	db := database.CallDB()
	// ---- db -------------- //
	// ----- router --------- //
	caleedroter := router.CallRouter(db)
	// ----- router  ----- //

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, caleedroter))
}
