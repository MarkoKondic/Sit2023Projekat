package main

import (
	"net/http"
	"ren-pl-backend/routes"
)

func main() {
	router := routes.SetupRouter()
	http.ListenAndServe(":8080", router)
}
