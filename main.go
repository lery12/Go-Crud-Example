package main

import (
	"crudExample/routes"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:8000", routes.GetRoutes())
}
