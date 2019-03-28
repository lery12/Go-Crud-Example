package routes

import (
	"goji.io"
	"goji.io/pat"
	"crudExample/controllers"
)

func GetRoutes() *goji.Mux {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/vehicle"), controllers.GetVehicleList)
	mux.HandleFunc(pat.Get("/vehicle/:id"), controllers.GetVehicle)
	mux.HandleFunc(pat.Post("/vehicle"), controllers.CreateVehicle)
	mux.HandleFunc(pat.Put("/vehicle/:id"), controllers.UpdateVehicle)
	mux.HandleFunc(pat.Delete("/vehicle/:id"), controllers.DeleteVehicle)
	return mux
}