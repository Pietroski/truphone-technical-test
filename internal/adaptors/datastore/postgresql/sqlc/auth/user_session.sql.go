// Code generated by sqlc. DO NOT EDIT.
// source: user_session.sql

package auth_store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUserSession = `-- name: CreateUserSession :one
INSERT INTO user_session (user_id, access_token, role, permissions, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING user_id, access_token, role, permissions, created_at, updated_at
`

type CreateUserSessionParams struct {
	UserID      uuid.UUID   `json:"userID"`
	AccessToken string      `json:"accessToken"`
	Role        Roles       `json:"role"`
	Permissions Permissions `json:"permissions"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func (q *Queries) CreateUserSession(ctx context.Context, arg CreateUserSessionParams) (UserSession, error) {
	row := q.queryRow(ctx, q.createUserSessionStmt, createUserSession,
		arg.UserID,
		arg.AccessToken,
		arg.Role,
		arg.Permissions,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i UserSession
	err := row.Scan(
		&i.UserID,
		&i.AccessToken,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserSessionByAccessToken = `-- name: DeleteUserSessionByAccessToken :exec
DELETE
FROM user_session
WHERE access_token = $1
`

func (q *Queries) DeleteUserSessionByAccessToken(ctx context.Context, accessToken string) error {
	_, err := q.exec(ctx, q.deleteUserSessionByAccessTokenStmt, deleteUserSessionByAccessToken, accessToken)
	return err
}

const deleteUserSessionByID = `-- name: DeleteUserSessionByID :exec
DELETE
FROM user_session
WHERE user_id = $1
`

func (q *Queries) DeleteUserSessionByID(ctx context.Context, userID uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteUserSessionByIDStmt, deleteUserSessionByID, userID)
	return err
}

const getUserSessionByAccessToken = `-- name: GetUserSessionByAccessToken :one
SELECT user_id, access_token, role, permissions, created_at, updated_at
FROM user_session
WHERE access_token = $1
LIMIT 1
`

func (q *Queries) GetUserSessionByAccessToken(ctx context.Context, accessToken string) (UserSession, error) {
	row := q.queryRow(ctx, q.getUserSessionByAccessTokenStmt, getUserSessionByAccessToken, accessToken)
	var i UserSession
	err := row.Scan(
		&i.UserID,
		&i.AccessToken,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserSessionByID = `-- name: GetUserSessionByID :one
SELECT user_id, access_token, role, permissions, created_at, updated_at
FROM user_session
WHERE user_id = $1
LIMIT 1
`

func (q *Queries) GetUserSessionByID(ctx context.Context, userID uuid.UUID) (UserSession, error) {
	row := q.queryRow(ctx, q.getUserSessionByIDStmt, getUserSessionByID, userID)
	var i UserSession
	err := row.Scan(
		&i.UserID,
		&i.AccessToken,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserSessionByUserAccessToken = `-- name: UpdateUserSessionByUserAccessToken :one
UPDATE user_session
SET (role, permissions, updated_at) = ($2, $3, $4)
WHERE access_token = $1
RETURNING user_id, access_token, role, permissions, created_at, updated_at
`

type UpdateUserSessionByUserAccessTokenParams struct {
	AccessToken string      `json:"accessToken"`
	Role        Roles       `json:"role"`
	Permissions Permissions `json:"permissions"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func (q *Queries) UpdateUserSessionByUserAccessToken(ctx context.Context, arg UpdateUserSessionByUserAccessTokenParams) (UserSession, error) {
	row := q.queryRow(ctx, q.updateUserSessionByUserAccessTokenStmt, updateUserSessionByUserAccessToken,
		arg.AccessToken,
		arg.Role,
		arg.Permissions,
		arg.UpdatedAt,
	)
	var i UserSession
	err := row.Scan(
		&i.UserID,
		&i.AccessToken,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserSessionByUserID = `-- name: UpdateUserSessionByUserID :one
UPDATE user_session
SET (access_token, role, permissions, updated_at) = ($2, $3, $4, $5)
WHERE user_id = $1
RETURNING user_id, access_token, role, permissions, created_at, updated_at
`

type UpdateUserSessionByUserIDParams struct {
	UserID      uuid.UUID   `json:"userID"`
	AccessToken string      `json:"accessToken"`
	Role        Roles       `json:"role"`
	Permissions Permissions `json:"permissions"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func (q *Queries) UpdateUserSessionByUserID(ctx context.Context, arg UpdateUserSessionByUserIDParams) (UserSession, error) {
	row := q.queryRow(ctx, q.updateUserSessionByUserIDStmt, updateUserSessionByUserID,
		arg.UserID,
		arg.AccessToken,
		arg.Role,
		arg.Permissions,
		arg.UpdatedAt,
	)
	var i UserSession
	err := row.Scan(
		&i.UserID,
		&i.AccessToken,
		&i.Role,
		&i.Permissions,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
