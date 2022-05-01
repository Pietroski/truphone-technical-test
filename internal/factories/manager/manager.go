package manager_factory

import (
	device_store "github.com/Pietroski/truphone-technical-test/internal/adaptors/datastore/postgresql/sqlc/manager/devices"
	manager_controller "github.com/Pietroski/truphone-technical-test/internal/controllers/manager"
	devices_factory "github.com/Pietroski/truphone-technical-test/internal/factories/manager/devices"
	"github.com/gin-gonic/gin"
)

type ManagerServer struct {
	address string
	router  *gin.Engine

	deviceStore device_store.Store
}

func NewManagerServer(deviceStore device_store.Store) *ManagerServer {
	// TODO: apply validations for arguments if needed

	factory := &ManagerServer{
		deviceStore: deviceStore,
	}

	factory.Handle()

	return factory
}

func (f *ManagerServer) Handle() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// TODO: implement error handling for CORS
	_ = r.SetTrustedProxies([]string{"localhost", "127.0.0.1"})

	managerController := manager_controller.NewManagerController(f.deviceStore)
	deviceFactory := devices_factory.NewDeviceFactory(managerController.Devices)

	v1 := r.Group("/v1")
	{
		manager := v1.Group("/manager")
		{
			deviceFactory.Handle(manager)
		}
	}

	f.router = r
}

func (f *ManagerServer) Start() error {
	// TODO: remove this after implemented via envs
	f.address = ":8089"
	return f.router.Run(f.address)
}
