package device_controller

import (
	device_store "github.com/Pietroski/truphone-technical-test/internal/adaptors/datastore/postgresql/sqlc/manager/devices"
	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	store device_store.Store
}

func NewDeviceController(store device_store.Store) *DeviceController {
	// TODO: apply validations for arguments if needed

	controller := &DeviceController{
		store: store,
	}

	return controller
}

func (c *DeviceController) CreateDevice(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) GetDevice(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) getDeviceByID(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) GettUserDevices(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) getUserDevices(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) getPaginatedUserDevices(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) ListDevices(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) listDevices(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) listPaginatedDevices(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) DeleteDevice(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) deleteDeviceByID(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) DeleteUserDevices(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) UpdateDevice(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *DeviceController) UpdateUserDevices(ctx *gin.Context) {
	// TODO: implement me!!
}
