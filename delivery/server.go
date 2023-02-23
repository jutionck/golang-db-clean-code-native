package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-clean-code-native-query/config"
	"github.com/jutionck/golang-clean-code-native-query/delivery/controller"
	"github.com/jutionck/golang-clean-code-native-query/repository"
	"github.com/jutionck/golang-clean-code-native-query/usecase"
)

type Server struct {
	vehicleUC usecase.VehicleUseCase
	engine    *gin.Engine
	host      string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	controller.NewVehicleController(s.engine, s.vehicleUC)
}

func NewServer() *Server {
	c := config.NewConfig()
	r := gin.Default()
	dbConn := config.NewDbConnection(c)
	vehicleRepo := repository.NewVehicleRepository(dbConn.Conn())
	vehilceUC := usecase.NewVehicleUseCase(vehicleRepo)
	if c.ApiHost == "" || c.ApiPort == "" {
		panic("No Host or port define")
	}
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &Server{vehicleUC: vehilceUC, engine: r, host: host}
}
