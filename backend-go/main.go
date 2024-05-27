package main

import (
	"backend-api/internal/controller"
	"backend-api/internal/service"
	"log"
	"net/http"
)

func main() {
	meli := service.NewMeliAPI(&http.Client{})
	controller := controller.NewItemsController(meli)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/items", controller.SearchItems)
	mux.HandleFunc("/api/items/{id}", controller.GetItemByID)

	log.Println("Server listening on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
