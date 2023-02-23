package repository

import (
	"database/sql"

	"github.com/jutionck/golang-clean-code-native-query/model"
)

type VehicleRepository interface {
	Store(newVehicle *model.Vehilce) error
	List(page int, totalRows int) ([]model.Vehilce, error)
	Get(id string) (model.Vehilce, error)
	Update(oldVehicle *model.Vehilce) error
	Destroy(id string) error
}

type vehilceRepository struct {
	db *sql.DB
}

func (v *vehilceRepository) Store(newVehicle *model.Vehilce) error {
	sql := "INSERT INTO vehicle (id, brand, model, year, weight) VALUES ($1, $2, $3, $4, $5)"
	_, err := v.db.Exec(sql, newVehicle.Id, newVehicle.Brand, newVehicle.Model, newVehicle.Year, newVehicle.Weight)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehilceRepository) List(page int, totalRows int) ([]model.Vehilce, error) {
	sql := `SELECT id, brand, model, year, weight FROM vehicle LIMIT $1 OFFSET $2`
	limit := totalRows
	offset := limit * (page - 1)
	rows, err := v.db.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}

	var vehicle []model.Vehilce
	for rows.Next() {
		var vehilce model.Vehilce
		err := rows.Scan(&vehilce.Id, &vehilce.Brand, &vehilce.Model, &vehilce.Year, &vehilce.Weight)
		if err != nil {
			return nil, err
		}
		vehicle = append(vehicle, vehilce)
	}
	return vehicle, nil
}

func (v *vehilceRepository) Get(id string) (model.Vehilce, error) {
	sql := `SELECT id, brand, model, year, weight FROM vehicle WHERE id = $1`
	var vehilce model.Vehilce
	err := v.db.QueryRow(sql, id).Scan(&vehilce.Id, &vehilce.Brand, &vehilce.Model, &vehilce.Year, &vehilce.Weight)
	if err != nil {
		return model.Vehilce{}, err
	}
	return vehilce, nil
}

func (v *vehilceRepository) Update(oldVehicle *model.Vehilce) error {
	sql := "UPDATE vehicle set brand = $1, model = $2, year = $3, weight = $4 WHERE id = $5"
	_, err := v.db.Exec(sql, oldVehicle.Brand, oldVehicle.Model, oldVehicle.Year, oldVehicle.Weight, oldVehicle.Id)
	if err != nil {
		return err
	}

	return nil
}

func (v *vehilceRepository) Destroy(id string) error {
	sql := "DELETE FROM vehicle WHERE id = $1"
	_, err := v.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func NewVehicleRepository(db *sql.DB) VehicleRepository {
	return &vehilceRepository{db: db}
}
