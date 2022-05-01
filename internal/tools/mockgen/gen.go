package mock_generator

// internal - usually adaptors
//go:generate mockgen -package mockedUserStore -destination ../../../internal/adaptors/datastore/postgresql/sqlc/auth/user/mock/mockedUserStore.go github.com/Pietroski/truphone-technical-test/internal/adaptors/datastore/postgresql/sqlc/auth/user Store
//go:generate mockgen -package mockedDeviceStore -destination ../../../internal/adaptors/datastore/postgresql/sqlc/manager/device/mock/mockedDeviceStore.go github.com/Pietroski/truphone-technical-test/internal/adaptors/datastore/postgresql/sqlc/manager/device Store

// external - usually server clients
// wrong example -> //go:generate mockgen -package mockedUserStore -destination ../../../internal/adaptors/datastore/postgresql/sqlc/auth/user/mock/mockedUserStore.go github.com/Pietroski/truphone-technical-test/adaptors/services/datastore/postgresql/sqlc/auth/user/ Store
