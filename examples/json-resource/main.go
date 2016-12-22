package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/buffalo/examples/json-resource/actions"
	"github.com/markbates/going/defaults"
)

func main() {
	port := defaults.String(os.Getenv("PORT"), "3000")
	log.Printf("Starting json-resource on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}