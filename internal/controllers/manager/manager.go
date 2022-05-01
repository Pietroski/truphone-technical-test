package manager_controller

import (
	device_store "github.com/Pietroski/truphone-technical-test/internal/adaptors/datastore/postgresql/sqlc/manager/devices"
	device_controller "github.com/Pietroski/truphone-technical-test/internal/controllers/manager/devices"
)

type ManagerController struct {
	Devices *device_controller.DeviceController
}

func NewManagerController(store device_store.Store) *ManagerController {
	// TODO: apply validations for arguments if needed

	controller := &ManagerController{
		Devices: device_controller.NewDeviceController(store),
	}

	return controller
}
