package main

import (
	database "cooking-web-app/db"
	"cooking-web-app/services/api"
	"log"
	"net/http"
)

func main() {

	db := database.GetDB()

	main_router := http.NewServeMux()
	api.GetApi(db).MountOnServe(main_router, "/api")

	server := &http.Server{
		Addr:    ":8080",
		Handler: main_router,
	}

	log.Println("Server is running on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
