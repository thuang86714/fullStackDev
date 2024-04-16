package main

import (
	"backend/initializers"
	"log"
	"net/http"

	"backend/platform/authenticator"
	"backend/platform/router"
)

func init(){
	initializers.LoadEnvVariables()
}

func main() { 
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}