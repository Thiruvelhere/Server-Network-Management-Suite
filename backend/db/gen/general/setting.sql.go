// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: setting.sql

package generaldb

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, role, email, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING id, name, role, email
`

type CreateUserParams struct {
	Name         string `db:"name" json:"name"`
	Role         string `db:"role" json:"role"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
}

type CreateUserRow struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name"`
	Role  string    `db:"role" json:"role"`
	Email string    `db:"email" json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Role,
		arg.Email,
		arg.PasswordHash,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Role,
		&i.Email,
	)
	return i, err
}
