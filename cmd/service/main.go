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
	f := services.FoobarService{}
	server := otohttp.NewServer()
	generated.RegisterGreeterService(server, g)
	generated.RegisterFoobarService(server, f)
	http.Handle("/oto/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
