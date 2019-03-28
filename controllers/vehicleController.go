package controllers

import (
	"crudExample/model"
	"encoding/json"
	"fmt"
	"goji.io/pat"
	"net/http"
)

func GetVehicleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model.GetVehicleList())
	if err != nil {
		w.WriteHeader(400)
	}
}

func GetVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	paramID := pat.Param(r, "id")
	if paramID != "" {
		v, err := model.GetVehicleByLicense(paramID)

		if err != nil {
			w.WriteHeader(204)
		} else {
			err = json.NewEncoder(w).Encode(v)
			if err != nil {
				w.WriteHeader(400)
			}
		}
	}
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := model.Vehicle{}

	err := json.NewDecoder(r.Body).Decode(&v)

	if err != nil {
		w.WriteHeader(400)
	} else {
		addedVehicle, err := model.AddVehicle(&v)
		if err != nil {
			w.WriteHeader(406)
		}
		err = json.NewEncoder(w).Encode(addedVehicle)
		if err != nil {
			w.WriteHeader(400)
		}
	}
}

func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var v model.Vehicle

	err := json.NewDecoder(r.Body).Decode(&v)
	paramID := pat.Param(r, "id")
	if paramID == v.License {
		if err != nil {
			fmt.Println("Error in encoding things", &v)
		}

		v, err = model.UpdateVehicle(&v)
		if err != nil {
			w.WriteHeader(406)
		}

		err = json.NewEncoder(w).Encode(v)
		if err != nil {
			w.WriteHeader(400)
		}
	}
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var v model.Vehicle

	err := json.NewDecoder(r.Body).Decode(&v)
	paramID := pat.Param(r, "id")
	if paramID == v.License {
		if err != nil {
			fmt.Println("Error in encoding things", &v)
		}

		v, err = model.RemoveVehicle(&v)
		if err != nil {
			w.WriteHeader(406)
		}

		err = json.NewEncoder(w).Encode(v)
		if err != nil {
			w.WriteHeader(400)
		}
	}
}
