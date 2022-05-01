package main

import (
	"database/sql"
	"log"

	auth_store "github.com/Pietroski/truphone-technical-test/internal/adaptors/datastore/postgresql/sqlc/auth"
	device_store "github.com/Pietroski/truphone-technical-test/internal/adaptors/datastore/postgresql/sqlc/manager/devices"
	auth_factory "github.com/Pietroski/truphone-technical-test/internal/factories/auth"
	manager_factory "github.com/Pietroski/truphone-technical-test/internal/factories/manager"
)

var (
	stopServerSig = make(chan error)
)

func main() {
	// TODO: pass database conn
	authStore := auth_store.NewStore(&sql.DB{})
	authServer := auth_factory.NewAuthServer(authStore)

	// TODO: pass database conn
	deviceStore := device_store.NewStore(&sql.DB{})
	managerServer := manager_factory.NewManagerServer(deviceStore)

	{ // server initialisation grouping

		// We do not need waiting group here
		// select statement will hold the server running
		// until an error is returned during initialisation

		// auth server initialisation
		go func(stopServerSig chan error) {
			if err := authServer.Start(); err != nil {
				stopServerSig <- err
			}
		}(stopServerSig)

		// manager server initialisation
		go func(stopServerSig chan error) {
			if err := managerServer.Start(); err != nil {
				stopServerSig <- err
			}
		}(stopServerSig)
	}

	select {
	case err := <-stopServerSig:
		log.Println(err)
		return
	}
}
