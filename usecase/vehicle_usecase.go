package usecase

import (
	"github.com/jutionck/golang-clean-code-native-query/model"
	"github.com/jutionck/golang-clean-code-native-query/repository"
	"github.com/jutionck/golang-clean-code-native-query/utils"
)

type VehicleUseCase interface {
	RegisterNewVehicle(newVehicle *model.Vehilce) error
	FindAllVehicle(page int, totalRows int) ([]model.Vehilce, error)
	FindByIdVehilce(id string) (model.Vehilce, error)
	UpdateBehilce(oldVehicle *model.Vehilce) error
	DestroyVehicle(id string) error
}

type vehilceUseCase struct {
	repo repository.VehicleRepository
}

func (v *vehilceUseCase) RegisterNewVehicle(newVehicle *model.Vehilce) error {
	newVehicle.Id = utils.UuidGenerate()
	return v.repo.Store(newVehicle)
}

func (v *vehilceUseCase) FindAllVehicle(page int, totalRows int) ([]model.Vehilce, error) {
	if page == 0 {
		page = 1
	}

	if totalRows == 0 {
		totalRows = 5
	}

	return v.repo.List(page, totalRows)
}

func (v *vehilceUseCase) FindByIdVehilce(id string) (model.Vehilce, error) {
	return v.repo.Get(id)
}

func (v *vehilceUseCase) UpdateBehilce(oldVehicle *model.Vehilce) error {
	return v.repo.Update(oldVehicle)
}

func (v *vehilceUseCase) DestroyVehicle(id string) error {
	return v.repo.Destroy(id)
}

func NewVehicleUseCase(repo repository.VehicleRepository) VehicleUseCase {
	return &vehilceUseCase{repo: repo}
}
