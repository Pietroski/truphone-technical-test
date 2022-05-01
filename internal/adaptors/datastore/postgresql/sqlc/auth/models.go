// Code generated by sqlc. DO NOT EDIT.

package auth_store

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Permissions string

const (
	PermissionsREADWRITE Permissions = "READ;WRITE"
	PermissionsREAD      Permissions = "READ"
)

func (e *Permissions) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Permissions(s)
	case string:
		*e = Permissions(s)
	default:
		return fmt.Errorf("unsupported scan type for Permissions: %T", src)
	}
	return nil
}

type Roles string

const (
	RolesADMIN  Roles = "ADMIN"
	RolesCOMMON Roles = "COMMON"
)

func (e *Roles) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Roles(s)
	case string:
		*e = Roles(s)
	default:
		return fmt.Errorf("unsupported scan type for Roles: %T", src)
	}
	return nil
}

type Recovering struct {
	UserID       uuid.UUID `json:"userID"`
	RecoveryLink string    `json:"recoveryLink"`
	ExpiryDate   time.Time `json:"expiryDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserSession struct {
	UserID      uuid.UUID   `json:"userID"`
	AccessToken string      `json:"accessToken"`
	Role        Roles       `json:"role"`
	Permissions Permissions `json:"permissions"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Users struct {
	RowID          int64       `json:"rowID"`
	UserID         uuid.UUID   `json:"userID"`
	Name           string      `json:"name"`
	Email          string      `json:"email"`
	HashedPassword string      `json:"hashedPassword"`
	Role           Roles       `json:"role"`
	Permissions    Permissions `json:"permissions"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}