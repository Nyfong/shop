package main

import (
	"log"
	"net/http"

	_ "restAPI/docs"
	"restAPI/route"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Shop API
//	@version		1.0
//	@description	Shop API built with Go
//	@host			localhost:8080
//	@BasePath		/
func main() {

    http.HandleFunc("GET /shop", route.GetAllFromShop)
    http.HandleFunc("GET /shop/{id}", route.GetSingleProduct)

	http.HandleFunc("/docs/", httpSwagger.WrapHandler)

	log.Println("Server running on http://localhost:8080")
	log.Println("Swagger UI: http://localhost:8080/docs/index.html")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}