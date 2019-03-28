package model

import (
	"crudExample/data"
	"errors"
	"fmt"
)

type Vehicletype interface {
	GetVehicleList()
	GetVehicleByLicense()
	AddVehicle()
	RemoveVehicle()
}

type Vehicle struct {
	License string `json:"id"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
}

func (Vehicle) Error() string {
	panic("Woeps")
}

func GetVehicleList() []Vehicle {
	vehicles := []Vehicle{}
	rows, err := data.DatabaseConnection.Query("SELECT id, model, year FROM vehicle")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id string
		var model string
		var year int
		err = rows.Scan(&id, &model, &year)
		vehicles = append(vehicles, Vehicle{id, model, year})
	}

	return vehicles
}

func GetVehicleByLicense(license string) (vehicle Vehicle, err error) {
	foundVehicle := Vehicle{}
	stmt, err := data.DatabaseConnection.Prepare(`SELECT id, model, year FROM vehicle WHERE vehicle.Id = $1`)
	rows, err := stmt.Query(license)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var id string
		var model string
		var year int
		err = rows.Scan(&id, &model, &year)
		foundVehicle = Vehicle{id, model, year}
	}

	return foundVehicle, nil
}

func UpdateVehicle(v *Vehicle) (vehcile Vehicle, err error) {
	if v.License != "" {
		sqlStmt := "UPDATE vehicle SET (id, model, year) = ($1, $2, $3) WHERE vehicle.id = $4"
		_, err := data.DatabaseConnection.Exec(sqlStmt, v.License, v.Model, v.Year, v.License)
		if err != nil {
			fmt.Println(err)
			return Vehicle{}, errors.New("Update went wrong!")
		}
		return *v, nil
	}
	return Vehicle{}, errors.New("Update went wrong! No Licenseplate...")
}

func AddVehicle(v *Vehicle) (vehicle Vehicle, err error) {
	if v.License != "" {

		sqlStatement := `INSERT INTO vehicle (id, model, year) VALUES($1, $2, $3)`
		_, err := data.DatabaseConnection.Exec(sqlStatement, v.License, v.Model, v.Year)

		if err != nil {

			return Vehicle{}, errors.New("No vehicle added!")
		}

		return *v, nil
	} else {

		return Vehicle{}, errors.New("No empty license!")
	}
}

func RemoveVehicle(v *Vehicle) (vehicle Vehicle, err error) {
	sqlStmt := "DELETE FROM vehicle WHERE id = $1"
	_, err = data.DatabaseConnection.Exec(sqlStmt, v.License)

	if err != nil {
		return Vehicle{}, errors.New("No delete action was exec!")
	}

	return *v, nil
}
