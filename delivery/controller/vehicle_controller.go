package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-clean-code-native-query/model"
	"github.com/jutionck/golang-clean-code-native-query/usecase"
)

type VehicleController struct {
	router  *gin.Engine
	usecase usecase.VehicleUseCase
}

func (vc *VehicleController) getAllVehicle(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	vehicles, err := vc.usecase.FindAllVehicle(page, size)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve vehicle data"})
		return
	}
	ctx.JSON(http.StatusOK, vehicles)
}
func (vc *VehicleController) getVehicleById(ctx *gin.Context) {
	id := ctx.Param("id")
	vehicle, err := vc.usecase.FindByIdVehilce(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve vehicle data"})
		return
	}
	ctx.JSON(http.StatusOK, vehicle)
}

func (vc *VehicleController) registerVehicle(ctx *gin.Context) {
	var vehicle model.Vehilce
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := vc.usecase.RegisterNewVehicle(&vehicle); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vehicle)
}

func (vc *VehicleController) updateVehicle(ctx *gin.Context) {
	var vehicle model.Vehilce
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := vc.usecase.UpdateBehilce(&vehicle); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vehicle)
}

func (vc *VehicleController) destoryVehicle(ctx *gin.Context) {
	id := ctx.Param("id")
	err := vc.usecase.DestroyVehicle(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.String(http.StatusNoContent, "")
}

func NewVehicleController(r *gin.Engine, usecase usecase.VehicleUseCase) *VehicleController {
	controller := VehicleController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/vehicle", controller.getAllVehicle)
	r.GET("/vehicle/:id", controller.getVehicleById)
	r.POST("/vehicle", controller.registerVehicle)
	r.PUT("/vehicle", controller.updateVehicle)
	r.DELETE("/vehicle/:id", controller.destoryVehicle)
	return &controller
}
