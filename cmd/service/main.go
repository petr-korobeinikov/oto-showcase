package main

import (
	"log"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"

	"oto-showcase/generated"
	"oto-showcase/services"
)

func main() {
	g := services.GreeterService{}
	server := otohttp.NewServer()
	generated.RegisterGreeterService(server, g)
	http.Handle("/oto/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
