// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.12.0

package querytest

import (
	"database/sql"

	"github.com/google/uuid"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type User struct {
	Sub uuid.UUID
}